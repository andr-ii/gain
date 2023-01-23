package metrics

import "github.com/andr-ii/punchy/conf"

func Generate(ch chan conf.AppData) {
	responses := initResponses()
	statistics := initStatistics()
	latency := initLatency()

	go runProgress()

	for data := range ch {
		if data.Rps != nil {
			go statistics.setRps(*data.Rps)
		}

		if data.Response != nil {
			go responses.update(data.Response.Status)
			go statistics.setTotal()
			go latency.update(data.Response.Latency)
		}
	}
}
