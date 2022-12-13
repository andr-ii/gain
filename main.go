package main

import (
	"andr-ll/plt/plan"
	"andr-ll/plt/request"
	"context"
	"fmt"
	"time"
)

type RequestId string

var REQUEST_ID RequestId = "requestId"

func main() {
	plan := plan.Get()

	for i := uint8(1); i <= plan.Request.Amount; i++ {
		url := fmt.Sprintf("%v://%v/%v/%v", plan.Protocol, plan.Host, *plan.Path, i)

		go request.Perform(
			context.WithValue(context.Background(), REQUEST_ID, 1),
			plan.Method,
			url,
			nil,
		)
		time.Sleep(1000 * time.Millisecond)
	}
}
