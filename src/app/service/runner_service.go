package service

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/kaikeventura/cat-runner/src/app/client"
	"github.com/kaikeventura/cat-runner/src/app/model"
)

type RunnerService struct {
	httpClient client.HttpClient
}

func ConstructRunnerService(httpClient client.HttpClient) RunnerService {
	return RunnerService{httpClient}
}

func (service RunnerService) RunHttp(httpRunner model.HttpRunner) model.Response {
	url := buildUrl(httpRunner.Http)
	body := httpRunner.Http.Body.ContentText

	resp, time, error := service.httpClient.Post(url, httpRunner.Http.Headers, body)

	if error != nil {
		return model.Response{}
	}

	responseStatus := model.ResponseStatus{
		Description: resp.Status,
		Code:        resp.StatusCode,
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
	}

	request := resp.Request
	requestBody, _ := io.ReadAll(request.Body) // TODO não esta fazendo o parse
	requestData := model.RequestData{
		Method: request.Method,
		Url:    buildFullyUrlResquest(*request.URL),
		Header: request.Header,
		Body:   string(requestBody),
	}

	return model.Response{
		ResponseStatus: responseStatus,
		ResponseBody:   string(responseBody),
		Time:           time,
		RequestData:    requestData,
	}
}

func buildUrl(http model.Http) string {
	protocol := strings.ToLower(string(http.Protocol))
	hostname := http.Host
	port := buildPort(http.Port)
	path := http.Path
	queryParams := buildQueryParams(http.Parameters)

	return protocol + "://" + hostname + port + "/" + path + queryParams
}

func buildPort(port int) string {
	if port == 0 {
		return ""
	}
	return ":" + strconv.Itoa(port)
}

func buildQueryParams(queryParams []model.KVParam) string {
	if len(queryParams) == 0 {
		return ""
	}

	var params = ""
	for _, param := range queryParams {
		params += "?" + param.Key + "=" + param.Value
	}

	return params
}

func buildFullyUrlResquest(requestUrl url.URL) string {
	baseUrl := requestUrl.Scheme + "://" + requestUrl.Host + requestUrl.Path
	if requestUrl.RawQuery != "" {
		return "?" + requestUrl.RawQuery
	}

	return baseUrl
}
