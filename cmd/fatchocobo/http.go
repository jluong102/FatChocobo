package main

import (
	"net/http"
)

func MakeRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}

	return client.Do(request)
}

func SetJson(request *http.Request) {
	request.Header["Content-Type"] = []string{"application/json"}
}
