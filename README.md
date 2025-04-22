# Label Exporter

Looks for docker containers with the label format *dns.name=foo* and *dns.value=10.0.0.1*
and provides this information via the API endpoint */dns*.

## How to build

If you are using Just, then run:
```sh
just build
```
otherwise
```sh
go build -o label-exporter cmd/label-exporter/main.go
```

## How to run

```sh
./label-exporter
```

## Status / Version

0.0.1. This is a hobby app and will likely experience many breaking changes or no changes.
