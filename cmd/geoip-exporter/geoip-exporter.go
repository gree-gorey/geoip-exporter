package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sync"
	"time"
	"strings"

	"github.com/gree-gorey/geoip-exporter/pkg/conn"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	location = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "job_location",
			Help: "Location connections number",
		},
		[]string{"location"},
	)
)

func main() {
	addr := flag.String("web.listen-address", ":9300", "Address on which to expose metrics")
	interval := flag.Int("interval", 10, "Interval fo metrics collection in seconds")
	blackList := flag.String("blacklist", "104.31.10.172,104.31.11.172", "Addresses blacklist to filter out from results")
	debug := flag.Bool("debug", false, "Debug log level")
	flag.Parse()
	prometheus.MustRegister(location)
	http.Handle("/metrics", prometheus.Handler())
	go run(int(*interval), *debug, *blackList)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func run(interval int, debug bool, blackList string) {
	for {
		var wg sync.WaitGroup
		c := conn.Connections{}
		wg.Add(1)
		blackListArr := strings.Split(blackList, ",")
		b := make(map[string]bool)
		for _, key := range blackListArr {
			b[key] = true
		}
		p := conn.Params{UseWg: true, Wg: &wg, BlackList: b}
		go c.RunJob(&p)
		wg.Wait()
		if debug == true {
			ser, err := json.Marshal(c)
			if err != nil {
				log.Println(err)
			}
			log.Println(string(ser))
		}
		location.Reset()
		for code, number := range c.ConnectionsByCode {
			location.With(prometheus.Labels{"location": code}).Set(float64(number))
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
