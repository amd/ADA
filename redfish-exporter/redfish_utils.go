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
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
)

const (
	PeriodicRetryTime = 30
)

type RedfishServer struct {
	IP        string `json:"ip"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginType string `json:"loginType"`
	SlurmNode string `json:"slurmNode"`
}

type RedfishServersCommongConfig struct {
	HostSuffix string `json:"hostSuffix"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
}

type SubscriptionPayload struct {
	Destination         string                           `json:"Destination,omitempty"`
	EventTypes          []redfish.EventType              `json:"EventTypes,omitempty"`
	RegistryPrefixes    []string                         `json:"RegistryPrefixes,omitempty"`
	ResourceTypes       []string                         `json:"ResourceTypes,omitempty"`
	DeliveryRetryPolicy redfish.DeliveryRetryPolicy      `json:"DeliveryRetryPolicy,omitempty"`
	HTTPHeaders         map[string]string                `json:"HttpHeaders,omitempty"`
	Oem                 interface{}                      `json:"Oem,omitempty"`
	Protocol            redfish.EventDestinationProtocol `json:"Protocol,omitempty"`
	Context             string                           `json:"Context,omitempty"`
}

type RedfishSubsciptionFailedData struct {
	server  RedfishServer
	payload SubscriptionPayload
}

// Create a new connection to a redfish server
func getRedfishClient(server RedfishServer, tlsTimeout string) (*gofish.APIClient, error) {
	timeOut := 0
	if tlsTimeout != "" {
		t, err := strconv.Atoi(tlsTimeout)
		if err == nil {
			timeOut = t
		}
	}

	clientConfig := gofish.ClientConfig{
		Endpoint:            server.IP,
		Username:            server.Username,
		Password:            server.Password,
		Insecure:            true, // TODO Set Based on login type
		TLSHandshakeTimeout: timeOut,
	}

	c, err := gofish.Connect(clientConfig)
	if err != nil {
		log.Printf("Error connecting to redfish server %s: %v", server.IP, err)
		return nil, err
	}

	log.Printf("Successfully connected to redfish server %s", server.IP)
	return c, nil
}

// Create a subscription
func createSubscription(c *gofish.APIClient, server RedfishServer, subscriptionPayload SubscriptionPayload) (string, error) {
	// Get the event service
	eventService, err := c.Service.EventService()
	if err != nil {
		return "", fmt.Errorf("failed to get event service on server %s: %v", server.IP, err)
	}

	if err := deleteConflictingSubscriptions(c, server, subscriptionPayload); err != nil {
		return "", fmt.Errorf("failed to delete conflicting subscriptions: %v", err)
	}

	// Create the subscription based on the Redfish version
	if isV1_5() {
		return createV1_5Subscription(eventService, subscriptionPayload)
	}
	return createLegacySubscription(eventService, subscriptionPayload)
}

func isV1_5() bool {
	// TODO Logic to determine if Redfish server is <v1.5 or higher
	// We assume false until we get version info on the servers.
	return false
}

// Create V1.5 subscription
func createV1_5Subscription(eventService *redfish.EventService, SubscriptionPayload SubscriptionPayload) (string, error) {
	subscriptionURI, err := eventService.CreateEventSubscriptionInstance(
		SubscriptionPayload.Destination,
		SubscriptionPayload.RegistryPrefixes,
		SubscriptionPayload.ResourceTypes,
		SubscriptionPayload.HTTPHeaders,
		SubscriptionPayload.Protocol,
		SubscriptionPayload.Context,
		SubscriptionPayload.DeliveryRetryPolicy,
		SubscriptionPayload.Oem,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create v1.5 subscription: %w", err)
	}

	return subscriptionURI, nil
}

// Create legacy subscription
func createLegacySubscription(eventService *redfish.EventService, SubscriptionPayload SubscriptionPayload) (string, error) {
	subscriptionURI, err := eventService.CreateEventSubscription(
		SubscriptionPayload.Destination,
		SubscriptionPayload.EventTypes,
		SubscriptionPayload.HTTPHeaders,
		SubscriptionPayload.Protocol,
		SubscriptionPayload.Context,
		SubscriptionPayload.Oem,
	)

	if err != nil {
		return "", fmt.Errorf("failed to create legacy subscription: %w", err)
	}

	return subscriptionURI, nil
}

// Create subscriptions for all servers and return their URIs
// Rollback if any subscription attempt fails
func CreateSubscriptionsForAllServers(redfishServers []RedfishServer, subscriptionPayload SubscriptionPayload, subscriptionMap map[string]string, mu *sync.Mutex, tlsTimeout string) error {
	failedSubsChan := make(chan RedfishSubsciptionFailedData)
	for _, server := range redfishServers {
		go doSubscription(server, subscriptionPayload, subscriptionMap, mu, failedSubsChan, tlsTimeout)
	}

	go periodicSubscriptionRetry(failedSubsChan, subscriptionMap, mu, tlsTimeout)
	return nil
}

func periodicSubscriptionRetry(failedSubsChan chan RedfishSubsciptionFailedData, subscriptionMap map[string]string, mu *sync.Mutex, tlsTimeout string) {
	failedSubsMap := map[string]RedfishSubsciptionFailedData{}

	ticker := time.NewTicker(PeriodicRetryTime * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for ip, data := range failedSubsMap {
				log.Printf("Retrying subscription for: %v", ip)
				go doSubscription(data.server, data.payload, subscriptionMap, mu, failedSubsChan, tlsTimeout)
				delete(failedSubsMap, ip)
			}
		case data := <-failedSubsChan:
			failedSubsMap[data.server.IP] = data
		}
	}
}

func doSubscription(server RedfishServer, subscriptionPayload SubscriptionPayload, subscriptionMap map[string]string, mu *sync.Mutex, failedSubsChan chan RedfishSubsciptionFailedData, tlsTimeout string) {
	c, err := getRedfishClient(server, tlsTimeout)
	if err != nil {
		log.Printf("[error] failed to connect to server %s: %v", server.IP, err)
		failedSubsChan <- RedfishSubsciptionFailedData{server: server, payload: subscriptionPayload}
		return
	}
	defer c.Logout()

	subscriptionURI, err := createSubscription(c, server, subscriptionPayload)
	if err != nil {
		log.Printf("[error] subscription failed on server %s: %v", server.IP, err)
		failedSubsChan <- RedfishSubsciptionFailedData{server: server, payload: subscriptionPayload}
		return
	}
	mu.Lock()
	subscriptionMap[server.IP] = subscriptionURI
	mu.Unlock()
	log.Printf("Successfully created subscription on Redfish server %s: %s", server.IP, subscriptionURI)
}

// Delete all event subscriptions stored in the map
func DeleteSubscriptionsFromAllServers(redfishServers []RedfishServer, subscriptionMap map[string]string) {
	var wg sync.WaitGroup

	log.Println("Unsubscribing from servers...")

	for serverIP, subscriptionURI := range subscriptionMap {
		wg.Add(1)
		go func(serverIP, subscriptionURI string) {
			defer wg.Done()
			server := getServerInfo(redfishServers, serverIP)

			c, err := getRedfishClient(server, "")
			if err != nil {
				log.Printf("Failed to connect to server %s: %v", server.IP, err)
				return
			}
			defer c.Logout()

			if err := deleteSubscriptionFromServer(c, server, subscriptionURI); err != nil {
				log.Printf("Failed to delete event subscription on server %s: %v", server.IP, err)
			} else {
				log.Printf("Successfully deleted event subscription from server %s: %s", server.IP, subscriptionURI)
			}
		}(serverIP, subscriptionURI)
	}

	wg.Wait()
}

// Delete a subscription from a redfish server
func deleteSubscriptionFromServer(c *gofish.APIClient, server RedfishServer, subscriptionURI string) error {
	// Get the event service
	eventService, err := c.Service.EventService()
	if err != nil {
		return fmt.Errorf("failed to get event service on server %s: %v", server.IP, err)
	}

	// Attempt to delete the subscription
	if err := eventService.DeleteEventSubscription(subscriptionURI); err != nil {
		return fmt.Errorf("failed to delete event subscription on server %s: %v", server.IP, err)
	}

	return nil
}

// Unsubscribes/deletes conflicting subscriptions from the server
func deleteConflictingSubscriptions(c *gofish.APIClient, server RedfishServer, subscriptionPayload SubscriptionPayload) error {
	subscriptions, err := getServerSubscriptions(c, server)
	if err != nil {
		return err
	}
	for _, subscription := range subscriptions {
		if subscription.Destination == subscriptionPayload.Destination {
			err := deleteSubscriptionFromServer(c, server, subscription.ODataID)
			if err != nil {
				return fmt.Errorf("failed to delete event subscription %s on server %s: %v", subscription.ID, server.IP, err)
			}
			log.Printf("Successfully deleted overlapping event subscription %s from server %s", subscription.ID, server.IP)
		}
	}
	return nil
}

// Gets all subscriptions currently active on the given server
func getServerSubscriptions(c *gofish.APIClient, server RedfishServer) ([]*redfish.EventDestination, error) {
	// Get the event service
	eventService, err := c.Service.EventService()
	if err != nil {
		return nil, fmt.Errorf("failed to get event service on server %s: %v", server.IP, err)
	}

	subscriptions, err := eventService.GetEventSubscriptions()
	if err != nil {
		return nil, fmt.Errorf("failed to get event subscriptions on server %s: %v", server.IP, err)
	}
	return subscriptions, nil
}

// Retrieve the server's credentials from the config based on IP
func getServerInfo(redfishServers []RedfishServer, serverIP string) RedfishServer {
	for _, redfishServer := range redfishServers {
		if redfishServer.IP == serverIP {
			return redfishServer
		}
	}
	return RedfishServer{}
}

func getServerInfoByIP(redfishServers []RedfishServer, IP string) RedfishServer {
	for _, redfishServer := range redfishServers {
		ip := redfishServer.IP
		if IP == extractHost(ip) {
			return redfishServer
		}
	}
	return RedfishServer{}
}
