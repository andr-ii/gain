package request

import (
	"andr-ll/gain/conf"
	"math"
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
		if plan.RPS.Max != nil && RPS >= *plan.RPS.Max {
			continue
		}

		if plan.RPS.Max != nil {
			RPS = int(math.Min(float64(RPS+*plan.RPS.Step), float64(*plan.RPS.Max)))
		} else {
			RPS += *plan.RPS.Step
		}

		sendRps(ch, &RPS)
	}
}

func sendRps(ch chan conf.AppData, RPS *int) {
	ch <- conf.AppData{
		Rps: RPS,
	}
}
