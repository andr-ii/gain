package metrics

import "andr-ll/plt/conf"

func Generate(ch chan conf.AppData) {
	status := newStatus()

	go runProgress()

	for data := range ch {
		if data.Rps != nil {
			go status.setRps(*data.Rps)
		}

		if data.Response != nil {
			go status.update(data.Response.Status)
		}
	}
}
