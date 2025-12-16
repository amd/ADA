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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var EventReceivedCountMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "redfish_events_received_total",
		Help: "Total number of Redfish events received",
	},
	[]string{"ip", "slurm_node_name", "event_severity", "message_id"},
)

var EventProcessingTimeMetric = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "RedFishEvents_processing_time",
		Help: "Time taken to process events",
	},
	[]string{"SourceIP", "Severity"}, // Define the labels you want to use
)

// Number of nodes by monitoring status
var RedfishExporterStatus = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "redfish_exporter_monitoring_status",
		Help: "Node monitoring status by type: TotalNodes, MonitoredNodes, MonitorFailures, DrainedNodes",
	},
	[]string{"type"},
)

// Total Drain API Per EventSeverity Per Node Per DrainReason Per APIStatus(Success/Failure)
var SLURMAPIDrainCallsMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "slurm_drain_api_calls_total",
		Help: "Total number of SLURM Drain API calls",
	},
	[]string{"ip", "slurm_node_name", "event_severity", "drain_reason", "api_status"}, // Define the labels you want to use
)

func init() {
	// Register the counter with Prometheus's default registry
	prometheus.MustRegister(EventReceivedCountMetric)
	prometheus.MustRegister(SLURMAPIDrainCallsMetric)

	// Register the gauge with Prometheus's default registry
	prometheus.MustRegister(EventProcessingTimeMetric)
	prometheus.MustRegister(RedfishExporterStatus)
}
