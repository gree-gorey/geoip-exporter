package conn

import (
	"log"

	"github.com/gree-gorey/geoip-exporter/pkg/geo"
	"github.com/shirou/gopsutil/net"
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
		if (conn.Status == "ESTABLISHED") && (conn.Raddr.IP != "127.0.0.1") &&
			(conn.Raddr.IP != "104.31.10.172") && (conn.Raddr.IP != "104.31.11.172") {
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

}
