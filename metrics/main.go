package metrics

import (
	"andr-ll/plt/plan"
)

func Generate(ch chan ResponseData, rps chan uint16) {
	plan := plan.Get()
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
