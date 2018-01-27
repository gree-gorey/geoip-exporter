package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sync"
	"time"

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
	debug := flag.Bool("debug", false, "Debug log level")
	flag.Parse()
	prometheus.MustRegister(location)
	http.Handle("/metrics", prometheus.Handler())
	go run(int(*interval), *debug)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func run(interval int, debug bool) {
	for {
		var wg sync.WaitGroup
		c := conn.Connections{}
		wg.Add(1)
		p := conn.Params{UseWg: true, Wg: &wg}
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
