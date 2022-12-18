package conf

import (
	"time"
)

type AppData struct {
	Response *ResponseData
	Rps      *int
}

type ResponseData struct {
	Status        string
	Latency       time.Duration
	ContentLength int64
}
