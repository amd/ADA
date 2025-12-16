/**
 * Copyright 2024 Advanced Micro Devices, Inc.  All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
**/

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/nod-ai/ADA/redfish-exporter/metrics"
	"github.com/nod-ai/ADA/redfish-exporter/slurm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		targetFile          string
		subscriptionMapLock sync.Mutex // to guard access to the map
	)

	flag.StringVar(&targetFile, "target", "", "Path to the target file for host/slurm node names")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting Redfish Event Listener/Exporter")

	// Setup configuration
	AppConfig := setupConfig(targetFile)

	// Read monitoring-disabled servers list from file and store them in config
	// These won't be subscribed to until monitoring is enabled for them
	serversIP := readMonitoringDisabledServersFromFile()

	for _, ip := range serversIP {
		if _, ok := AppConfig.RedfishServers[ip]; ok {
			AppConfig.RedfishServers[ip].MonitoringDisabled = true
		}
	}

	// Log the initialized config
	log.Printf("Initialized Config: %+v", AppConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var slurmQueue *slurm.SlurmQueue

	slurmQueue = slurm.InitSlurmQueue(ctx)
	go slurmQueue.ProcessEventActionQueue()

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the listener
	listener := NewServer(AppConfig.SystemInformation.ListenerIP, AppConfig.SystemInformation.ListenerPort, slurmQueue)
	go func() {
		if err := listener.Start(AppConfig); err != nil {
			log.Printf("Server error: %v", err)
			close(sigChan)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Printf("Starting metrics server on :%d", AppConfig.SystemInformation.MetricsPort)
		portStr := strconv.Itoa(AppConfig.SystemInformation.MetricsPort)
		if err := http.ListenAndServe(":"+portStr, nil); err != nil && err != http.ErrServerClosed {
			log.Printf("Metrics server error: %v", err)
		}
	}()

	// Initialize Prometheus metrics with default values
	initMetrics()

	subscriptionMap := make(map[string]string)

	// Subscribe the listener to the event stream for all servers
	err := CreateSubscriptionsForAllServers(AppConfig.RedfishServers, AppConfig.SubscriptionPayload, subscriptionMap, &subscriptionMapLock, &AppConfig.RWMutex, AppConfig.TlsTimeOut)
	if err != nil {
		log.Fatal(err)
	}

	AppConfig.runHttpRequestsHandlerServer()

	// Wait for shutdown signal
	<-sigChan
	log.Println("Received shutdown signal. Closing listener...")

	// Signal shutdown to all components
	close(listener.shutdownChan)

	// Give some time for the listener to close gracefully
	time.Sleep(time.Second)

	// Unsubscribe the listener from all servers
	subscriptionMapLock.Lock()
	DeleteSubscriptionsFromAllServers(AppConfig.RedfishServers, subscriptionMap)
	subscriptionMapLock.Unlock()

	cancel()

	time.Sleep(time.Second)

	// Perform any additional shutdown steps here
	log.Println("Shutdown complete")
}

func initMetrics() {
	metrics.RedfishExporterStatus.WithLabelValues("TotalNodes").Set(float64(0))
	metrics.RedfishExporterStatus.WithLabelValues("MonitoredNodes").Set(float64(0))
	metrics.RedfishExporterStatus.WithLabelValues("MonitorFailures").Set(float64(0))
	metrics.RedfishExporterStatus.WithLabelValues("DrainedNodes").Set(float64(0))
}

// Helper function to extract the host from the URL
func extractHost(url string) string {
	// Remove http:// or https://
	host := strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://")

	// Trim the port if available
	if splitHost, _, err := net.SplitHostPort(host); err == nil {
		host = splitHost
	}
	return host
}

func (c *Config) runHttpRequestsHandlerServer() error {
	// start a http server
	router := mux.NewRouter()

	// API handler for fetching all the nodes where monitoring is enabled
	router.Methods(http.MethodGet).Subrouter().Handle("/api/nodes", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.RLock()
		servers := []string{}
		for ip, server := range c.RedfishServers {
			if !server.MonitoringDisabled {
				servers = append(servers, ip)
			}
		}
		c.RUnlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(servers)
	}))

	// API handler to enable/disable monitoring on a node
	router.Methods(http.MethodPatch).Subrouter().Handle("/api/nodes/{serverIP}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		serverIP, ok := vars["serverIP"]
		if !ok {
			http.Error(w, fmt.Sprintf("Server IP not specified"), http.StatusInternalServerError)
			return
		}

		var payload struct {
			MonitoringEnabled bool `json:"monitoring-enabled"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		c.RLock()
		defer c.RUnlock()

		if _, ok := c.RedfishServers[serverIP]; ok {
			if c.RedfishServers[serverIP].MonitoringDisabled != !payload.MonitoringEnabled {
				// For application restarts, persist the monitoring-disabled servers list
				servers := []string{}
				for ip, server := range c.RedfishServers {
					if server.MonitoringDisabled && server.IP != serverIP {
						servers = append(servers, ip)
					}
				}

				respString := fmt.Sprintf("node %s enabled for monitoring\n", serverIP)
				if !payload.MonitoringEnabled {
					respString = fmt.Sprintf("node %s disabled for monitoring\n", serverIP)
					servers = append(servers, serverIP)
				}
				totalNodesMonitored := len(c.RedfishServers) - len(servers)
				metrics.RedfishExporterStatus.WithLabelValues("TotalNodes").Set(float64(totalNodesMonitored))
				writeMonitoringDisabledServersToFile(servers)
				go ProcessMonitoringEnableDisableEvent(payload.MonitoringEnabled, c.RedfishServers[serverIP], c.SubscriptionPayload)

				w.Write([]byte(respString))
				return
			}
		}
		w.WriteHeader(http.StatusOK)
	}))

	// start the server
	go http.ListenAndServe(c.SystemInformation.RestURL, router)

	return nil
}
