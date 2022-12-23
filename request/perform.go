package request

import (
	"andr-ll/gain/conf"
	"net/http"
	"time"
)

func perform(ch chan conf.AppData) {
	client := http.Client{}
	request := makeRequest()

	startTime := time.Now()
	res, err := client.Do(request)

	var status string
	var content int64

	if err != nil || res == nil {
		status = "503 Service Unavailable"
		content = 0
	} else {
		status = res.Status
		content = res.ContentLength
	}

	ch <- conf.AppData{
		Response: &conf.ResponseData{
			Status:        status,
			Latency:       float32(time.Since(startTime).Seconds()),
			ContentLength: content,
		},
	}
}
