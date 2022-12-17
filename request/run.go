package request

import (
	"andr-ll/plt/conf"
	"andr-ll/plt/metrics"
	"time"
)

func Run(ch chan metrics.ResponseData, rps chan uint16) {
	plan := conf.Plan
	RPS := plan.RPS.Value

	rps <- RPS

	go start(&RPS, ch)

	for {
		<-time.After(time.Duration(*plan.RPS.Interval) * time.Second)
		RPS += *plan.RPS.Step
		rps <- RPS
	}
}
