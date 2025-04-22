[private]
default:
    @just --list

run:
    @go run cmd/label-exporter/main.go

build:
    @go build -o label-exporter cmd/label-exporter/main.go

add-container:
    docker run --rm -d --label dns.name=foo --label dns.value=10.0.0.1 busybox /bin/sh -c "sleep 300"