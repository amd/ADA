# Metrics Documentation

The AMD Redfish Exporter exposes the following Prometheus-compatible metrics:

## Event Metrics

### `redfish_events_received_total`

- **Type**: Counter
- **Description**: Total number of Redfish events received
- **Labels**:
  - `ip`: IP address of the Redfish server that sent the event
  - `slurm_node_name`: Node name of the SLURM node
  - `event_severity`: Severity of the Redfish event
  - `message_id`: Event message kind

### `RedFishEvents_processing_time`

- **Type**: Gauge
- **Description**: Time taken to process Redfish events in seconds
- **Labels**:
  - `SourceIP`: IP address of the Redfish server that sent the event
  - `Severity`: Severity of the Redfish event

### `redfish_exporter_monitoring_status`

- **Type**: Gauge
- **Description**: Node monitoring status by type.
  Valid Type Values:
  - TotalNodes: Number of nodes the application is supposed to be monitoring
  - MonitoredNodes: Number of nodes successfully being monitored
  - MonitorFailures: Number of nodes where monitoring has failed
  - DrainedNodes: Number of nodes that have been drained
- **Labels**:
  - `type`: Node Monitoring status types - TotalNodes, MonitoredNodes, MonitorFailures, DrainedNodes

### `slurm_drain_api_calls_total`

- **Type**: Counter
- **Description**: Total number of Slurm API calls that succeeded
- **Labels**:
  - `ip`: IP address of the Redfish server that sent the event
  - `slurm_node_name`: Node name of the SLURM node
  - `event_severity`: Severity of the event
  - `drain_reason`: SLURM Node drain reason
  - `api_status`: SLURM Node Drain API Status. Valid values: Success, Failure


## Accessing Metrics

Metrics are exposed on the `/metrics` endpoint. By default, this is available at `http://localhost:2112/metrics`.

## Prometheus Configuration

To scrape metrics from the AMD Redfish Exporter, add the following job to your Prometheus configuration:

```yaml
scrape_configs:
  - job_name: 'amd_redfish_exporter'
    static_configs:
      - targets: ['localhost:2112']
```

[Placeholder: Additional Prometheus config options?]
