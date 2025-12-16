# AMD Redfish Exporter

The AMD Redfish Exporter is a Go server that subscribes to events and telemetry from servers with Redfish service enabled. It exposes this data on a Prometheus-compatible metrics route for scraping. The exporter also integrates with SLURM to drain nodes when critical hardware events are detected. This exporter is tested on Slurm version `23.02.7` with REST API version `v0.0.39`

## Features

- Subscribes to Redfish events from multiple servers
- Processes and filters events based on configurable criteria
- Exposes Prometheus-compatible metrics
- Integrates with SLURM for node management
- Supports SSL/TLS for secure communication

## Installation

### Prerequisites

- Go 1.23 or later
- Docker
- SLURM (for node management functionality)

### Building from source

- Clone the repository:

```bash
git clone https://github.com/nod-ai/ADA.git
cd ADA/redfish-exporter
```

- Build the project:

```bash
make build
```

- The binary will be created as `amd-redfish-exporter` in the current directory.

### Using Docker

A Dockerfile is provided to build and run the AMD Redfish Exporter in a container:

- Build the Docker image:

```bash
docker build -t amd-redfish-exporter .
```

- Run the container:

```bash
docker run --env-file .env -p 8080:8080 -p 2112:2112 amd-redfish-exporter
```

Note: The `--env-file` flag takes a filename as an argument and expects each line to be in the `VAR=VAL` format. The `.env` file variables should not be wrapped in quotes. Example:

```bash
SUBSCRIPTION_PAYLOAD={"Destination": "http://host.docker.internal:8080", "EventTypes": ["Alert", "StatusChange"], "Protocol": "Redfish", "Context": "YourContextData"}
REDFISH_SERVERS=[{"ip": "http://<ip>:<port>", "username": "<username>", "password": "<password>", "loginType": "Session", "slurmNode": "Node1"}]
```

## Configuration

The AMD Redfish Exporter is configured using environment variables or a configuration file. Here are the main configuration options:

- `LISTENER_IP`: IP address to listen on (default: "127.0.0.1")
- `LISTENER_PORT`: Port to listen on for Redfish events (default: "8080")
- `METRICS_PORT`: Port to expose Prometheus metrics (default: "2112")
- `USE_SSL`: Enable SSL/TLS (default: "false")
- `CERTFILE`: Path to SSL certificate file (required if USE_SSL is true)
- `KEYFILE`: Path to SSL key file (required if USE_SSL is true)
- `SLURM_USER`: SLURM user name. Please note that only the `root` user has permission to perform node operations, such as draining.
- `SLURM_TOKEN`: SLURM authentication token
- `SLURM_CONTROL_NODE`: SLURM control node address
- `REDFISH_SERVERS`: JSON array of Redfish server configurations
- `SUBSCRIPTION_PAYLOAD`: JSON object defining the Redfish subscription parameters

For a complete list of configuration options, refer to the [Configuration Guide](docs/configuration.md).

## Usage

- Set up the configuration as described in the [Configuration](#configuration) section.

- Run the AMD Redfish Exporter:

```bash
./amd-redfish-exporter
```

- The exporter will start subscribing to events from the configured Redfish servers and expose metrics on the specified metrics port.

- Configure your Prometheus server to scrape this endpoint: `http://[server IP]:2112/metrics`.

Note that SSL can be enabled using the provided dev certificate and key and `--ssl` flag. For more detailed usage instructions, see the [User Guide](docs/user-guide.md).

### Slurm Integration

The exporter is designed to utilize Slurm APIs from version `0.0.39`. Please follow these instructions only if you wish to use APIs from a different version. Keep in mind that the code generator may not function properly with older versions. If you prefer to stick with the `0.0.39` APIs, you can skip steps 1 and 2.

1. Follow [this](./api/README.md) guide for generating the Slurm REST API.

2. Required configurations:

- **SLURM_TOKEN**: Obtain the token by running the following command on the Slurm control node (where slurmrestd is running):

```sh
scontrol token username=root lifespan=18000
```

(`lifespan` is specified in seconds)

- **SLURM_CONTROL_NODE**: The hostname or IP address of the Slurm control node where the REST server is running.

3. Run the AMD Redfish Exporter with `--enable-slurm` flag to enable SLURM integration:

```bash
./amd-redfish-exporter --enable-slurm
```

## Metrics

The exporter provides the following metrics:

- `redfish_exporter_monitoring_status`: Counter of Redfish events received per SLURM node per message kind per event severity
- `RedFishEvents_processing_time`: Gauge of event processing time
- `redfish_exporter_monitoring_status`: Gauge of node monitoring status by type: TotalNodes, MonitoredNodes, MonitorFailures, DrainedNodes
- `slurm_drain_api_calls_total`: Counter of SLURM Drain API calls per event severity per SLURM Node per drain reason per API Status(Success/Failure)

For a complete list of available metrics and their descriptions, see the [Metrics Documentation](docs/metrics.md).

## Testing

See [testing.md](docs/testing.md) for information on using the Mock Server and Integration Testing.

## Support

If you encounter any issues or have questions, please file an issue.
