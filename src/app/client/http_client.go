package client

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kaikeventura/cat-runner/src/app/model"
)

type HttpClient struct{}

func ConstructHttpClient() HttpClient {
	return HttpClient{}
}

func (httpClient HttpClient) Post(url string, headers []model.KVParam, body string) (*http.Response, int, error) {
	requestBody := strings.NewReader(body)
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			request.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(request)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	return resp, int(endTime.Sub(startTime).Milliseconds()), nil
}
