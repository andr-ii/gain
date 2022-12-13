package request

import (
	"io"
	"net/http"
)

func readResponse(resp *http.Response) string {
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)

	if err != nil {
		panic("An error ocurred while reading")
	}

	return string(resBody[:])
}
