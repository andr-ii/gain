package request

import (
	"andr-ll/plt/metrics"
	"net/http"
	"time"
)

func perform(ch chan metrics.ResponseData) {
	client := http.Client{}
	request := makeRequest()

	startTime := time.Now()
	resp, err := client.Do(request)

	if err != nil {
		panic("An error ocurred during the request")
	}

	ch <- metrics.ResponseData{
		Status:        resp.Status,
		Latency:       time.Since(startTime),
		ContentLength: uint64(resp.ContentLength),
	}
}
