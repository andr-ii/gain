package conf

import (
	"time"
)

type AppData struct {
	Response *ResponseData
	Rps      *uint16
}

type ResponseData struct {
	Status        string
	Latency       time.Duration
	ContentLength uint64
}
