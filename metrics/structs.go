package metrics

import "sync"

type response struct {
	id          uint8
	amount      uint32
	set         bool
	labelLength int
}

type status struct {
	mut       *sync.Mutex
	responses map[string]response
	total     uint32
	rps       uint16
}
