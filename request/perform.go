package request

import (
	"net/http"
)

func Perform(ch chan string, method, url string, bodyStruct *interface{}) {
	request := makeRequest(method, url, bodyStruct)

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		panic("An error ocurred during the request")
	}

	ch <- resp.Status
}
