# GeoIP Prometheus exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/gree-gorey/geoip-exporter)](https://goreportcard.com/report/github.com/gree-gorey/geoip-exporter)

`github.com/gree-gorey/geoip-exporter`

GeoIP exporter collects metrics about TCP-connections, 
locates remote IP-address and exposes metrics to Prometheus 
via `/metrics` endpoint.

Example visualization using Grafana:
![map](https://raw.githubusercontent.com/gree-gorey/geoip-exporter/master/static/map.png "map")

### Usage

Available command-line options:
```bash
--debug
    Debug log level
--interval int
    Interval fo metrics collection in seconds (default 10)
--web.listen-address string
    Address on which to expose metrics (default ":9300")
```

### Example

Example usage:
```bash
$ ./geoip-exporter --interval=10 --web.listen-address=127.0.0.1:9400 --debug
```



### TODO

- [X] Add command-line flags
- [ ] Add filter for IP addresses (blacklist)
