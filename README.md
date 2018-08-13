# GeoIP Prometheus exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/gree-gorey/geoip-exporter)](https://goreportcard.com/report/github.com/gree-gorey/geoip-exporter)  
[![](https://images.microbadger.com/badges/image/greegorey/geoip-exporter.svg)](https://microbadger.com/images/greegorey/geoip-exporter "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/greegorey/geoip-exporter.svg)](https://microbadger.com/images/greegorey/geoip-exporter "Get your own version badge on microbadger.com")

`go get github.com/gree-gorey/geoip-exporter`

GeoIP exporter collects metrics about TCP-connections, 
locates remote IP-address and exposes metrics to Prometheus 
via `/metrics` endpoint.

Example visualization using Grafana:  

![map](https://raw.githubusercontent.com/gree-gorey/geoip-exporter/master/static/map.png "map")

### Usage

Available command-line options:
```console
--blacklist string
    	Addresses blacklist to filter out from results (default "104.31.10.172,104.31.11.172")
--debug
    Debug log level
--interval int
    Interval fo metrics collection in seconds (default 10)
--web.listen-address string
    Address on which to expose metrics (default ":9300")
```

### Example

Example usage:
```console
$ ./geoip-exporter --interval=10 --web.listen-address=127.0.0.1:9400 --blacklist="104.31.10.172,104.31.11.172" --debug
```

## Quick start guide with Grafana

### 1. Run geoip-exporter

Download latest release:
```console
$ cd /tmp
$ curl -s https://api.github.com/repos/gree-gorey/geoip-exporter/releases/latest \
| grep "browser_download_url" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -qi - -O geoip-exporter
$ chmod +x geoip-exporter
$ mv geoip-exporter /usr/local/bin
```

Create service:
```console
# cat << GEO > /etc/systemd/system/geoip-exporter.service
[Unit]
Description=Geo IP exporter for Prometheus
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
Restart=Always
ExecStart=/usr/local/bin/geoip-exporter --interval=300 --web.listen-address=127.0.0.1:9300 --blacklist="104.31.10.172,104.31.11.172"

[Install]
WantedBy=multi-user.target
GEO
# systemctl enable geoip-exporter.service
# systemctl start geoip-exporter.service
```

Check that the service is running and send the responce:
```console
$ netstat -plnt | grep 9300
tcp        0      0 127.0.0.1:9300        0.0.0.0:*               LISTEN      2156/geoip-exporter
$ curl -s 127.0.0.1:9300/metrics | grep ^job_location
job_location{location="US"} 1
```

### 2. Tell Prometheus to collect metrics from geoip-exporter

Change Prometheus configuration file and add the following lines:
```
  - job_name: 'GeoIPExporter'
    scrape_interval: 10s
    static_configs:
      - targets: ['127.0.0.1:9300']
```

Then reload Proetheus:
```console
# pgrep "^prometheus$" | xargs -i kill -HUP {}
```

Go to Prometheus UI and check that it collects metrics:  

![map](https://raw.githubusercontent.com/gree-gorey/geoip-exporter/master/static/prom.png "map")


### 3. Setup Grafana plugin

Then you need to install Worldmap Panel plugin for Grafana:
```console
# grafana-cli plugins install grafana-worldmap-panel
```

Go to Grafana UI and add new panel (add panel -> Worldmap Panel).   
Go to the *Metrics* tab and add query:
```
sum(job_location) by (location)
```
Legend format:
```
{{location}}
```
Mark checkpoint *Instant*.  
Your setting should look like this:  

![map](https://raw.githubusercontent.com/gree-gorey/geoip-exporter/master/static/wm1.png "map")

Then go to the *Worldmap* tab and set up it as this:  

![map](https://raw.githubusercontent.com/gree-gorey/geoip-exporter/master/static/wm2.png "map")

From this point on your worldmap panel is ready.

### TODO

- [X] Add command-line flags
- [x] Add filter for IP addresses (blacklist)
- [X] Docker image
