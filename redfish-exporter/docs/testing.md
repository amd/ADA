# Mock Server and Integration Testing Documentation

## Running the Mock Redfish Server

The mock Redfish server simulates a Redfish-enabled device, allowing you to test the AMD Redfish Exporter without connecting to real hardware.

### Using Docker

- Run the following command to start the mock Redfish server:

```bash
docker run -d --net="host" --name redfish-mock dmtf/redfish-mockup-server:latest --port=8000
```

Note: The `--net="host"` option allows the container to access the host network directly. This option is necessary when the listener app is running directly on the host.

- Verify that the server is running:

```bash
curl http://localhost:8000/redfish/v1/
```

You should receive a JSON response representing the root of the Redfish service.

### Using Kubernetes

- Apply the Kubernetes configuration:

```bash
kubectl apply -k dev/redfish-mock
```

- Forward the service port to your local machine:

```bash
kubectl port-forward -n silo services/redfish-mock 8000:8000
```

- Verify that the server is running:

```bash
curl http://localhost:8000/redfish/v1/
```

#### Testing a Subscription

To manually create a subscription, create a JSON file (e.g. `data.json`) with the following content:

```json
{
    "Destination":"http://localhost:8080/",
    "Protocol":"Redfish",
    "EventFormatType":"Event",
    "IncludeOriginOfCondition":"true",
    "EventTypes":[
       "Alert",
       "ResourceRemoved",
       "ResourceAdded",
       "ResourceUpdated",
       "StatusChange"
    ]
 }
```

Send a POST request to register this subscription:

```bash
curl -d "@data.json" -X POST http://localhost:8000/redfish/v1/EventService/Subscriptions
```

#### Verifying a Subscription

Verify the subscription is successfully created:

```bash
curl http://localhost:8000/redfish/v1/EventService/Subscriptions
curl http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

#### Trigger a Test Event

You can trigger a test event by creating a file (e.g., `testevent.json`) with the following content:

```json
{
    "EventType": "Alert",
    "EventId": "TestEvent",
    "EventTimestamp": "2024-09-30T10:10:10Z",
    "Severity": "Warning",
    "Message": "This is a test event message",
    "MessageId": "Basic.1.5.SomethingIsHappening",
    "MessageArgs": ["1", "slot 3"],
    "OriginOfCondition": "/redfish/v1/Systems/1/Storage"
}
```

Submit the test event with:

```bash
curl -d "@testevent.json" -X POST http://localhost:8000/redfish/v1/EventService/Actions/EventService.SubmitTestEvent
```

> [!WARNING]
> Unlike what's shown [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/main/public-rackmount1/EventService/SubmitTestEventActionInfo/index.json) all the fields are required by the mockup server. otherwise, the request will result in a BadRequest 400. check [here](https://github.com/DMTF/Redfish-Mockup-Server/blob/2d39eb14122337ceab0712a9610b1cd37c65f487/redfishMockupServer.py#L169) for more info.

#### Delete a Subscription

To delete a subscription, use its ID in a DELETE request:

```bash
curl -X DELETE http://localhost:8000/redfish/v1/EventService/Subscriptions/[ID]
```

## Integration Testing

Integration tests ensure that the AMD Redfish Exporter works correctly with the mock Redfish server and other components.

### Prerequisites

- Go 1.23 or later
- Docker
- Make

### Running the Integration Tests

1. Navigate to the project root directory.

2. Run the integration tests using the following command:

```bash
make integration-test
```

This command will:

- Build the AMD Redfish Exporter
- Start two mock Redfish servers using Docker
- Run the integration tests
- Clean up the test environment
- Monitor the output for test results and any error messages.

### Test Coverage

The integration tests cover the following scenarios:

1. Redfish event subscription
2. Event reception and processing
3. Metric exposition

## Troubleshooting

If you encounter issues while running the mock server or integration tests, try the following:

- Ensure all prerequisites are installed and up to date.
- Check that no other services are using the required ports (8000 for the mock server, 8443 for the exporter).
- Review the Docker and Kubernetes logs for any error messages:

```bash
# For Docker
docker logs redfish-mock

# For Kubernetes
kubectl logs -n silo deployment/redfish-mock
```

- Verify that your firewall or security settings are not blocking the required network connections. If problems persist, please file an issue on the project's GitHub repository with detailed information about the error and your environment.
