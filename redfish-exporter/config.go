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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const (
	DefaultListenerPort   = "8080"
	DefaultMetricsPort    = "2112"
	DefaultUseSSL         = "false"
	DefaultSeverityConfig = "Fatal,Critical,Informational"
	NodeDrainPolicyFile   = "nodeDrainPolicy.json"
)

type Config struct {
	Information struct {
		Updated     string
		Description string
	}
	SystemInformation struct {
		ListenerIP   string
		ListenerPort string
		UseSSL       bool
		MetricsPort  int
	}
	CertificateDetails struct {
		CertFile string
		KeyFile  string
	}
	SlurmToken           string
	SlurmControlNode     string
	SlurmUser            string
	SlurmScontrolPath    string
	SlurmDrainExcludeStr string
	SubscriptionPayload  SubscriptionPayload
	RedfishServers       []RedfishServer
	TriggerEvents        map[string]map[string][]EventInfo //map[Severity][MessageRegistry.MessageId][]EventInfo
	PrometheusConfig     PrometheusConfig
	context              *tls.Config
	eventCount           int
	dataBuffer           []byte
	TlsTimeOut           string
}

type EventInfo struct {
	UniqueString      string
	Category          string
	Subcategory       string
	DrainReasonPrefix string
}

type TriggerEvent struct {
	Severity          string `json:"Severity"`
	Action            string `json:"Action"`
	Message           string `json:"Message"`
	DrainReasonPrefix string `json:"DrainReasonPrefix"`
}

type TriggerEventsInfo struct {
	Category          string `json:"Category"`
	Subcategory       string `json:"Subcategory"`
	MessageRegistry   string `json:"MessageRegistry"`
	MessageId         string `json:"MessageId"`
	UniqueString      string `json:"UniqueString"`
	Severity          string `json:"Severity"`
	DrainReasonPrefix string `json:"DrainReasonPrefix"`
	Enable            bool   `json:"Enable"`
}

type PrometheusConfig struct {
	Severity []string `json:"Severity"`
}

type target struct {
	Targets []string          `yaml:"targets"`
	Labels  map[string]string `yaml:"labels"`
}

func setupConfig(targetFile string) Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file. Assuming environment variables are set.")
	}

	// Initialize Config object
	var AppConfig Config

	// Read environment variables
	AppConfig.Information.Updated = os.Getenv("UPDATED")
	AppConfig.Information.Description = os.Getenv("DESCRIPTION")
	AppConfig.SystemInformation.ListenerIP = os.Getenv("LISTENER_IP")

	// Read and parse LISTENER_PORT with a default value
	listenerPort := os.Getenv("LISTENER_PORT")
	if listenerPort == "" {
		listenerPort = DefaultListenerPort
	}
	AppConfig.SystemInformation.ListenerPort = listenerPort

	// Metrics Port Configuration
	metricsPortStr := os.Getenv("METRICS_PORT")
	if metricsPortStr == "" {
		metricsPortStr = DefaultMetricsPort
	}
	metricsPort, err := strconv.Atoi(metricsPortStr)
	if err != nil {
		log.Fatalf("Failed to parse METRICS_PORT: %v", err)
	}
	AppConfig.SystemInformation.MetricsPort = metricsPort
	// Read and parse USE_SSL with a default value
	useSSLStr := os.Getenv("USE_SSL")
	if useSSLStr == "" {
		useSSLStr = DefaultUseSSL
	}
	useSSL, err := strconv.ParseBool(useSSLStr)
	if err != nil {
		log.Fatalf("Failed to parse USE_SSL: %v", err)
	}
	AppConfig.SystemInformation.UseSSL = useSSL

	AppConfig.CertificateDetails.CertFile = os.Getenv("CERTFILE")
	AppConfig.CertificateDetails.KeyFile = os.Getenv("KEYFILE")

	AppConfig.SlurmToken = os.Getenv("SLURM_TOKEN")
	AppConfig.SlurmControlNode = os.Getenv("SLURM_CONTROL_NODE")
	AppConfig.SlurmUser = os.Getenv("SLURM_USER")
	AppConfig.SlurmDrainExcludeStr = os.Getenv("SLURM_DRAIN_EXCLUDE_REASON_LIST")
	AppConfig.SlurmScontrolPath = os.Getenv("SLURM_SCONTROL_PATH")
	AppConfig.TlsTimeOut = os.Getenv("TLS_TIMEOUT")

	subscriptionPayloadJSON := os.Getenv("SUBSCRIPTION_PAYLOAD")
	if err := json.Unmarshal([]byte(subscriptionPayloadJSON), &AppConfig.SubscriptionPayload); err != nil {
		log.Fatalf("Failed to parse SUBSCRIPTION_PAYLOAD: %v", err)
	}

	prometheusConfigJSON := os.Getenv("PROMETHEUS_CONFIG")
	if prometheusConfigJSON != "" {
		err = json.Unmarshal([]byte(prometheusConfigJSON), &AppConfig.PrometheusConfig)
		if err != nil {
			log.Fatalf("Failed to unmarshal PROMETHEUS_CONFIG: %v", err)
		}
	}
	if prometheusConfigJSON == "" {
		AppConfig.PrometheusConfig.Severity = strings.Split(DefaultSeverityConfig, ",")
	}

	// Read and parse the REDFISH_SERVERS environment variable
	redfishServersJSON := os.Getenv("REDFISH_SERVERS")
	if redfishServersJSON == "" {
		log.Println("REDFISH_SERVERS environment variable is not set or is empty")
	} else {
		if err := json.Unmarshal([]byte(redfishServersJSON), &AppConfig.RedfishServers); err != nil {
			log.Fatalf("Failed to parse REDFISH_SERVERS: %v", err)
		}
	}

	// Read the node drain policy config file
	nodeDrainPolicyConfig, err := os.ReadFile(NodeDrainPolicyFile)

	if err != nil {
		log.Fatalf("Failed to read: %v", NodeDrainPolicyFile)
	}

	triggerEventsInfo := []TriggerEventsInfo{}
	err = json.Unmarshal(nodeDrainPolicyConfig, &triggerEventsInfo)
	if err != nil {
		log.Fatalf("Failed to unmarshal file: %v | err: %v", NodeDrainPolicyFile, err)
	}

	tInfoMap := map[string]map[string][]EventInfo{}

	for _, evt := range triggerEventsInfo {
		fmt.Printf("Trigger Event: %+v\n", evt)
		if evt.Enable != true {
			continue
		}
		eInfo := EventInfo{}
		eInfo.Category = evt.Category
		eInfo.Subcategory = evt.Subcategory
		eInfo.DrainReasonPrefix = evt.DrainReasonPrefix
		eInfo.UniqueString = evt.UniqueString
		key := ""
		if evt.MessageRegistry == "" {
			key = evt.MessageId
		} else {
			key = evt.MessageRegistry + "." + evt.MessageId
		}
		if ee, ok := tInfoMap[evt.Severity]; !ok {
			eInfoMap := map[string][]EventInfo{}
			eInfoMap[key] = []EventInfo{eInfo}
			tInfoMap[evt.Severity] = eInfoMap
		} else {
			ee[key] = append(ee[key], eInfo)
		}
	}

	AppConfig.TriggerEvents = tInfoMap

	for kk, tt := range AppConfig.TriggerEvents {
		fmt.Println("Severity: ", kk)
		for kkk, ttt := range tt {
			fmt.Println("key: ", kkk)
			fmt.Printf("event: %+v\n", ttt)
		}
	}

	// Read and parse the REDFISH_SERVERS_COMMON_CONFIG environment variable
	redfishServersCommonConfigJSON := os.Getenv("REDFISH_SERVERS_COMMON_CONFIG")
	if redfishServersCommonConfigJSON == "" {
		log.Println("redfishServersCommonConfigJSON environment variable is not set or is empty")
		return AppConfig
	}
	redfishServersCommonConfig := RedfishServersCommongConfig{}
	if err := json.Unmarshal([]byte(redfishServersCommonConfigJSON), &redfishServersCommonConfig); err != nil {
		log.Fatalf("Failed to parse REDFISH_SERVERS_COMMON_CONFIG: %v", err)
	}

	if targetFile == "" {
		log.Println("No target file provided")
		return AppConfig
	}

	targetYamlFile, err := os.ReadFile(targetFile)

	if err != nil {
		log.Fatalf("Failed to read file: %v", targetFile)
	}

	targets := []target{}

	err = yaml.Unmarshal(targetYamlFile, &targets)

	if err != nil {
		log.Fatalf("Error parsing target file: %v | err: %v", targetFile, err)
	}

	for _, t := range targets {
		log.Println("target: ", t.Targets)

		for _, hostName := range t.Targets {
			// add this target to Redfish servers
			server := RedfishServer{}
			bmcHost := fmt.Sprintf(hostName+".%v", redfishServersCommonConfig.HostSuffix)
			ips, err := net.LookupIP(bmcHost)
			if err != nil || len(ips) == 0 {
				log.Printf("[error] Couldn't get the IP for host: %v | ips: %v | err: %v", bmcHost, ips, err)
				continue
			}
			log.Println("IPs: ", ips)

			server.IP = fmt.Sprintf("https://%v", ips[0])
			server.LoginType = "Session"
			server.Username = redfishServersCommonConfig.UserName
			server.Password = redfishServersCommonConfig.Password
			server.SlurmNode = hostName
			AppConfig.RedfishServers = append(AppConfig.RedfishServers, server)
		}
	}

	return AppConfig
}
