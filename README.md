# GeoIP Prometheus exporter

`go get github.com/domainr/whois`

GeoIP exporter collects metrics about TCP-connections, 
locates remote IP-address and exposes metrics to Prometheus 
via `/metrics` endpoint.

### TODO

- [X] Add command-line flags
- [ ] Add filter for IP addresses (blacklist)
