package metrics

import "time"

type coordinates struct {
	row    uint8
	col    uint8
	width  uint8
	height uint8
}

type ResponseData struct {
	Status        string
	Latency       time.Duration
	ContentLength uint64
}
