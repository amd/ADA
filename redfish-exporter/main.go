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
	"flag"
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

	// Log the initialized config
	log.Printf("Initialized Config: %+v", AppConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var slurmQueue *slurm.SlurmQueue

	slurmQueue = slurm.InitSlurmQueue(ctx)
	go slurmQueue.ProcessEventActionQueue()

	subscriptionMap := make(map[string]string)

	// Subscribe the listener to the event stream for all servers
	err := CreateSubscriptionsForAllServers(AppConfig.RedfishServers, AppConfig.SubscriptionPayload, subscriptionMap, &subscriptionMapLock, AppConfig.TlsTimeOut)
	if err != nil {
		log.Fatal(err)
	}

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
	initMetrics(AppConfig.RedfishServers, AppConfig.PrometheusConfig.Severity)

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

func initMetrics(redfishServers []RedfishServer, severities []string) {
	for _, server := range redfishServers {
		host := extractHost(server.IP)

		// Initialize counters for each severity label
		// Additional label combinations should be defined here
		for _, severity := range severities {
			metrics.EventCountMetric.WithLabelValues(host, severity).Add(0)
		}
	}
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
