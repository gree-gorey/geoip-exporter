package main

import (
	"github.com/gree-gorey/geoip-exporter/pkg/conn"
	"sync"
)

func main() {
	
	var wg sync.WaitGroup
	
	c := conn.Connections{}

	wg.Add(1)
	
	p := conn.Params{UseWg: true, Wg: &wg}
	
	go c.RunJob(&p)
	
	wg.Wait()
}
