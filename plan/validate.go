package plan

import (
	"fmt"
	"net/http"
)

var methods = [4]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

func validateMethod(inMethod *string) {
	for _, method := range methods {
		if method == *inMethod {
			return
		}
	}

	panic(fmt.Sprintf("HTTP method is not allowed: %v\n", *inMethod))
}

func validateIntervalAndStep(plan *Plan) {
	interval := plan.Request.RPS.Interval
	step := plan.Request.RPS.Step

	if interval == nil {
		newInterval := plan.Duration * 60
		plan.Request.RPS.Interval = &newInterval
	}

	if step == nil {
		newStep := uint16(0)
		plan.Request.RPS.Step = &newStep
	}
}
