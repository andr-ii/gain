package metrics

import "sync"

type response struct {
	id          int
	amount      int
	set         bool
	labelLength int
}

type status struct {
	mut       *sync.Mutex
	responses map[string]response
	total     int
}
