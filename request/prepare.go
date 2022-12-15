package request

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func prepare(method, url string, body *interface{}) *http.Request {
	var reader io.Reader

	if body == nil {
		reader = bytes.NewReader([]byte{})
	} else {
		reqBody, err := json.Marshal(*body)

		if err != nil {
			panic("Could not convert to json")
		}

		reader = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequestWithContext(context.Background(), method, url, reader)

	if err != nil {
		panic("Could not create a request")
	}

	req.Header.Add("Content-Type", "application/json")

	return req
}
