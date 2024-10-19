#!/bin/bash

echo "Starting Backend Server 1..."
go run ./cmd/backend/backend1/main.go &

echo "Starting Backend Server 2..."
go run ./cmd/backend/backend2/main.go &

echo "Starting Proxy Server..."
go run ./cmd/proxy/main.go

wait
