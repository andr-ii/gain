package metrics

import (
	"andr-ll/plt/conf"
)

func Generate(ch chan ResponseData, rps chan uint16) {
	plan := conf.Plan
	status := newStatus()

	go newProgress(plan.Duration)

	for {
		select {
		case data := <-ch:
			go status.update(data.Status)
		case amount := <-rps:
			go status.setRps(amount)
		}
	}
}
