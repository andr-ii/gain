package request

import (
	"andr-ll/plt/metrics"
	"net/http"
	"time"
)

func perform(ch chan metrics.ResponseData, method, url string, bodyStruct *interface{}) {
	client := http.Client{}
	request := prepare(method, url, bodyStruct)

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
