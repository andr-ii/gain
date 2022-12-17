package request

import (
	"andr-ll/plt/conf"
	"net/http"
	"time"
)

func perform(ch chan conf.AppData) {
	client := http.Client{}
	request := makeRequest()

	startTime := time.Now()
	resp, err := client.Do(request)

	if err != nil {
		panic("An error ocurred during the request")
	}

	ch <- conf.AppData{
		Response: &conf.ResponseData{
			Status:        resp.Status,
			Latency:       time.Since(startTime),
			ContentLength: uint64(resp.ContentLength),
		},
	}
}
