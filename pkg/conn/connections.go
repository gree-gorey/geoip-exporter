package conn

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/net"
	"log"
	"github.com/gree-gorey/geoip-exporter/pkg/geo"
)

type Connections struct {
	ConnectionsByCode map[string]int `json:"connections_by_code"`
}

func (c *Connections) RunJob(p *Params) {
	if p.UseWg {
		defer p.Wg.Done()
	}
	c.GetActiveConnections()
}

func (c *Connections) GetActiveConnections() {

	cs, err := net.Connections("tcp")
	if err != nil {
		log.Println(err)
	}
	
	c.ConnectionsByCode = make(map[string]int)
	for _, conn := range cs {
		if (conn.Status == "ESTABLISHED") && (conn.Raddr.IP != "127.0.0.1") {
			code := geo.GetCode(conn.Raddr.IP)
			if code != "" {
				_, ok := c.ConnectionsByCode[code]
				if ok == true {
					c.ConnectionsByCode[code] += 1
				} else {
					c.ConnectionsByCode[code] = 1
				}
			}
		}

	}
	
	ser, err := json.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(ser))
	
}
