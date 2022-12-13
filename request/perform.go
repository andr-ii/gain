package request

import (
	"context"
	"fmt"
	"net/http"
)

func Perform(ctx context.Context, method, url string, bodyStruct *interface{}) string {
	request := makeRequest(ctx, method, url, bodyStruct)

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic("An error ocurred during the request")
	}

	res := readResponse(resp)

	fmt.Printf("%v\n", res)

	return res
}
