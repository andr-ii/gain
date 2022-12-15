package request

import (
	"andr-ll/plt/metrics"
	"andr-ll/plt/plan"
	"time"
)

func Run(ch chan metrics.ResponseData, rps chan uint16, plan plan.Plan) {
	RPS := plan.Request.RPS.Value
	interval := *plan.Request.RPS.Interval
	step := *plan.Request.RPS.Step

	rps <- RPS

	go start(&RPS, ch, plan)

	for {
		<-time.After(time.Duration(interval) * time.Second)
		RPS += step
		rps <- RPS
	}
}
