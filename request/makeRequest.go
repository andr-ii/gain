package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeRequest(ctx context.Context, method, url string, body *interface{}) *http.Request {
	validateMethod(&method)

	reader := bodyToReader(body)
	req, err := http.NewRequestWithContext(ctx, method, url, reader)

	if err != nil {
		panic("Could not create a request")
	}

	req.Header.Add("Content-Type", "application/json")

	return req
}

func bodyToReader(body *interface{}) io.Reader {
	if body == nil {
		return bytes.NewReader([]byte{})
	}

	reqBody, err := json.Marshal(*body)

	if err != nil {
		panic("Could not convert to json")
	}

	return bytes.NewReader(reqBody)
}

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
