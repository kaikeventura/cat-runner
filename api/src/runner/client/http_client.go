package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/kaikeventura/cat-runner/src/runner/model"
)

type HttpClient struct{}

func ConstructHttpClient() HttpClient {
	return HttpClient{}
}

var handler = map[model.HttpMethod]func(url string, headers []model.KVParam, body *string) (*http.Response, int, error){
	model.GET:    ConstructHttpClient().get,
	model.POST:   ConstructHttpClient().post,
	model.PUT:    ConstructHttpClient().put,
	model.PATCH:  ConstructHttpClient().patch,
	model.DELETE: ConstructHttpClient().delete,
}

func (httpClient HttpClient) MakeRequest(method model.HttpMethod, url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	handler, exists := handler[method]
	if exists {
		return handler(url, headers, body)
	} else {
		fmt.Println("Http method no mapped")
	}
	return nil, 0, nil
}

func (httpClient HttpClient) get(url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	requester, err := buildRequester("GET", url, body)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			requester.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(requester)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	time := int(endTime.Sub(startTime).Milliseconds())

	return resp, time, nil
}

func (httpClient HttpClient) post(url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	requester, err := buildRequester("POST", url, body)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			requester.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(requester)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	time := int(endTime.Sub(startTime).Milliseconds())

	return resp, time, nil
}

func (httpClient HttpClient) put(url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	requester, err := buildRequester("PUT", url, body)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			requester.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(requester)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	time := int(endTime.Sub(startTime).Milliseconds())

	return resp, time, nil
}

func (httpClient HttpClient) patch(url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	requester, err := buildRequester("PATCH", url, body)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			requester.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(requester)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	time := int(endTime.Sub(startTime).Milliseconds())

	return resp, time, nil
}

func (httpClient HttpClient) delete(url string, headers []model.KVParam, body *string) (*http.Response, int, error) {
	requester, err := buildRequester("DELETE", url, body)
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return nil, 0, err
	}

	if len(headers) > 0 {
		for _, header := range headers {
			requester.Header.Set(header.Key, header.Value)
		}
	}

	client := &http.Client{}

	startTime := time.Now()
	resp, err := client.Do(requester)
	endTime := time.Now()

	if err != nil {
		fmt.Println("Error sending the request:", err)
		return nil, 0, err
	}

	time := int(endTime.Sub(startTime).Milliseconds())

	return resp, time, nil
}

func buildRequester(methoid string, url string, body *string) (*http.Request, error) {
	var requestBody io.Reader

	if body != nil {
		requestBody = strings.NewReader(*body)
	}

	return http.NewRequest(methoid, url, requestBody)
}
