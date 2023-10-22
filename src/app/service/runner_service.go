package service

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kaikeventura/cat-runner/src/app/client"
	"github.com/kaikeventura/cat-runner/src/app/model"
)

type RunnerService struct {
	httpClient client.HttpClient
}

func ConstructRunnerService(httpClient client.HttpClient) RunnerService {
	return RunnerService{httpClient}
}

var allResponses = []model.Response{}

func (service RunnerService) Run(httpRunner model.HttpRunner) []model.Response {
	var wg sync.WaitGroup
	virtualUser := httpRunner.VirtualUser

	for user := 1; user <= virtualUser.UsersAmount; user++ {
		wg.Add(1)
		userRequester := virtualUserRequester{
			UserId:             user,
			InteractionsAmount: virtualUser.InteractionsAmount,
			InteractionDelay:   virtualUser.InteractionDelay,
		}
		go service.runHttp(httpRunner.Http, userRequester, &wg)
	}

	wg.Wait()

	return allResponses
}

func (service RunnerService) runHttp(http model.Http, virtualUser virtualUserRequester, wg *sync.WaitGroup) {
	defer wg.Done()

	url := buildUrl(http)
	body := http.Body.ContentText

	for requestIdx := 1; requestIdx <= virtualUser.InteractionsAmount; requestIdx++ {
		time.Sleep(time.Duration(virtualUser.InteractionDelay) * time.Millisecond)

		resp, time, error := service.httpClient.MakeRequest(http.HttpMethod, url, http.Headers, &body)

		if error != nil {
			return
		}

		responseStatus := model.ResponseStatus{
			Description: resp.Status,
			Code:        resp.StatusCode,
		}

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler o corpo da resposta:", err)
			return
		}

		request := resp.Request
		requestBody, _ := io.ReadAll(request.Body) // TODO não esta fazendo o parse
		requestData := model.RequestData{
			Method: request.Method,
			Url:    buildFullyUrlResquest(*request.URL),
			Header: request.Header,
			Body:   string(requestBody),
		}

		allResponses = append(
			allResponses,
			model.Response{
				ResponseStatus: responseStatus,
				ResponseBody:   string(responseBody),
				Time:           time,
				VirtualUserResponse: model.VirtualUserResponse{
					UserId:    virtualUser.UserId,
					RequestId: requestIdx,
				},
				RequestData: requestData,
			},
		)
	}
}

type virtualUserRequester struct {
	UserId             int
	InteractionsAmount int
	InteractionDelay   int
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
		return baseUrl + "?" + requestUrl.RawQuery
	}

	return baseUrl
}
