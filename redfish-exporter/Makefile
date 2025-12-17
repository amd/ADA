BINARY_NAME=amd-redfish-exporter

.PHONY: build run test clean

build:
	go build -o $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

integration-test: build
	./integration_test.sh
