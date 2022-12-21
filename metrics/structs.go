package metrics

import "sync"

type response struct {
	id          int
	amount      int
	set         bool
	labelLength int
}

type responses struct {
	mut       *sync.Mutex
	responses map[string]response
}

type statistics struct {
	mut   *sync.Mutex
	rps   int
	total int
}

type latency struct {
	mut    *sync.Mutex
	min    float32
	max    float32
	avrg   float32
	sum    float32
	amount float32
}
