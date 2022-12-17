package request

import (
	"andr-ll/plt/metrics"
	"time"
)

func start(RPS *uint16, ch chan metrics.ResponseData) {
	for {
		for i := 0; i < int(*RPS); i++ {
			go perform(ch)
		}

		time.Sleep(time.Second)
	}
}
