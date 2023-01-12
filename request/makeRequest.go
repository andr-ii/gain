package request

import (
	"andr-ii/gain/conf"
	"andr-ii/gain/random"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func makeRequest() *http.Request {
	var reader io.Reader
	plan := conf.Plan

	if plan.Body.Value == nil {
		reader = bytes.NewReader([]byte{})
	} else {
		reader = generateBody(&plan)
	}

	req, err := http.NewRequestWithContext(context.Background(), plan.Method, plan.Url, reader)

	if err != nil {
		panic("Could not create a request")
	}

	req.Header.Add("Content-Type", "application/json")

	return req
}

func generateBody(plan *conf.PlanEntity) io.Reader {
	body := make(map[string]any)

	for key, val := range *plan.Body.Value {
		body[key] = val
	}

	for _, opts := range *plan.Body.DynFields {
		if opts.Type == "string" && opts.Length != nil {
			body[opts.Key] = random.Str(*opts.Length)
			continue
		}

		if opts.Type == "number" && opts.Range != nil {
			nums := *opts.Range
			body[opts.Key] = random.Num(nums[0], nums[1])
			continue
		}

		if opts.Values != nil {
			values := *opts.Values
			id := random.Num(0, len(*opts.Values))

			body[opts.Key] = values[id]
			continue
		}
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		panic("Could not convert body to json")
	}

	return bytes.NewReader(jsonBody)
}
