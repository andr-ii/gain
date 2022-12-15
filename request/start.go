package request

import (
	"andr-ll/plt/metrics"
	"andr-ll/plt/plan"
	"time"
)

func start(RPS *uint16, ch chan metrics.ResponseData, plan plan.Plan) {
	for {
		for i := 0; i < int(*RPS); i++ {
			go perform(ch, plan.Method, plan.Url, plan.Request.Body)
		}

		time.Sleep(time.Second)
	}
}
