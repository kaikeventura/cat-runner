package client

import (
	"fmt"
	"net/http"

	"github.com/kaikeventura/cat-runner/src/app/model"
)

type HttpClient struct{}

func ConstructHttpClient() HttpClient {
	return HttpClient{}
}

func (httpClient HttpClient) Post(url string, headers []model.KVParam, body *string) *http.Response {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil
	}

	if len(headers) > 0 {
		for _, header := range headers {
			request.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil
	}
	defer resp.Body.Close()

	return resp
}
