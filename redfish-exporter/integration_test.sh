
#!/bin/bash
set -e

GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

log() { echo -e "${BLUE}[$(date '+%H:%M:%S')]${NC} $1"; }
success() { echo -e "  ${GREEN}✓${NC} $1"; }

log "Starting integration test setup..."

cleanup() {
    echo "Cleaning up..."

    if [ ! -z "$EXPORTER_PID" ]; then
        echo "Shutting down AMD Redfish Exporter..."
        kill -SIGINT $EXPORTER_PID
        wait $EXPORTER_PID 2>/dev/null
    fi

    echo "Stopping and removing Docker containers..."
    docker stop integration-test1 integration-test2 || true
    docker rm integration-test1 integration-test2 || true

    # Remove generated certificates
    rm -rf ./certs

    # Remove exporter binary
    rm -f amd-redfish-exporter
}
trap cleanup EXIT

docker rm -f integration-test1 integration-test2 2>/dev/null || true

log "Starting mock Redfish servers..."
docker run -d --name integration-test1 -p 8000:8000 dmtf/redfish-mockup-server:latest --port=8000 >/dev/null
docker run -d --name integration-test2 -p 8001:8000 dmtf/redfish-mockup-server:latest --port=8000 >/dev/null

log "Waiting for mock servers to be ready..."
for port in 8000 8001; do
    for i in {1..30}; do
        if curl -s http://localhost:$port/redfish/v1/ >/dev/null; then
            success "Mock server on port $port is ready"
            break
        fi
        [ $i -eq 30 ] && { echo "Error: Mock server on port $port failed to start"; exit 1; }
        sleep 1
    done
done

log "Verifying Redfish root endpoints..."
for port in 8000 8001; do
    if curl -s http://localhost:$port/redfish/v1/ | grep -q '"@odata.id": "/redfish/v1/"'; then
        success "Server on port $port has a valid Redfish root"
    else
        echo "Error: Server on port $port does not have a valid Redfish root"
        exit 1
    fi
done

log "Building AMD Redfish Exporter..."
go build -o amd-redfish-exporter

# Generate self-signed certificate for testing TLS
CERT_DIR="./certs"
mkdir -p $CERT_DIR
CERT_FILE="$CERT_DIR/server.crt"
KEY_FILE="$CERT_DIR/server.key"

if [ ! -f "$CERT_FILE" ] || [ ! -f "$KEY_FILE" ]; then
    log "Generating self-signed certificate for TLS..."
    openssl req -x509 -newkey rsa:2048 -keyout $KEY_FILE -out $CERT_FILE -days 1 -nodes -subj "/CN=localhost"
fi

log "Starting AMD Redfish Exporter with TLS..."
LISTENER_IP=0.0.0.0 \
LISTENER_PORT=8443 \
USE_SSL=true \
CERTFILE=$CERT_FILE \
KEYFILE=$KEY_FILE \
SERVERS='[{"ip":"http://localhost:8000","username":"root","password":"password","loginType":"basic"},{"ip":"http://localhost:8001","username":"root","password":"password","loginType":"basic"}]' \
SUBSCRIPTION_PAYLOAD='{"Destination":"https://localhost:8443","EventTypes":["Alert","ResourceRemoved","ResourceAdded","ResourceUpdated","StatusChange"],"Context":"IntegrationTest","Protocol":"Redfish"}' \
./amd-redfish-exporter &

EXPORTER_PID=$!

log "Waiting for AMD Redfish Exporter to be ready..."
for i in {1..30}; do
    if curl -s http://localhost:2112/metrics >/dev/null; then
        success "AMD Redfish Exporter is ready"
        break
    fi
    [ $i -eq 30 ] && { echo "Error: AMD Redfish Exporter failed to start"; exit 1; }
    sleep 1
done

log "Sending a test event to the AMD Redfish Exporter over HTTPS..."
curl -k -s -X POST https://localhost:8443 \
     -H "Content-Type: application/json" \
     -d '{
         "Events": [{
             "EventType": "Alert",
             "EventId": "TestEvent",
             "EventTimestamp": "2023-09-30T12:00:00Z",
             "Severity": "OK",
             "Message": "Test event",
             "MessageId": "TestEvent.1.0.TestMessage",
             "Context": "IntegrationTest"
         }]
     }' >/dev/null

log "Waiting for metrics to update..."
sleep 5

log "Checking Prometheus metrics..."
metrics=$(curl -s http://localhost:2112/metrics)

check_metric() {
    if echo "$metrics" | grep -q "$1"; then
        success "$1 metric is present"
        echo "    $(echo "$metrics" | grep "$1")"
    else
        echo "Error: $1 metric is missing"
        exit 1
    fi
}

check_metric "RedFishEvents_recieved"
check_metric "RedFishEvents_processing_time"
check_metric "SlurmAPI_failure"
check_metric "SlurmAPI_success"

log "Integration test completed successfully"

