package request

import (
	"andr-ll/plt/conf"
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
		reqBody, err := json.Marshal(*plan.Body.Value)

		if err != nil {
			panic("Could not convert to json")
		}

		reader = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequestWithContext(context.Background(), plan.Method, plan.Url, reader)

	if err != nil {
		panic("Could not create a request")
	}

	req.Header.Add("Content-Type", "application/json")

	return req
}
