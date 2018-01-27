// Package conn implements function for collecting
// active TCP connections.

package conn

import "sync"

// Type Params stores parameters.
type Params struct {
	UseWg bool
	Wg    *sync.WaitGroup
}
