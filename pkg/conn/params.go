package conn

import "sync"

type Params struct {
	UseWg bool
	Wg    *sync.WaitGroup
}
