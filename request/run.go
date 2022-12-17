package request

import (
	"andr-ll/plt/conf"
	"time"
)

func Run(ch chan conf.AppData) {
	plan := conf.Plan
	RPS := plan.RPS.Value

	sendRps(ch, &RPS)

	go func() {
		for {
			for i := 0; i < int(RPS); i++ {
				go perform(ch)
			}

			time.Sleep(time.Second)
		}
	}()

	for {
		<-time.After(time.Duration(*plan.RPS.Interval) * time.Second)
		RPS += *plan.RPS.Step
		sendRps(ch, &RPS)
	}
}

func sendRps(ch chan conf.AppData, RPS *uint16) {
	ch <- conf.AppData{
		Rps: RPS,
	}
}
