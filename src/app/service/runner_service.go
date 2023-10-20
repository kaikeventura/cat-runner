package service

import (
	"strconv"
	"strings"

	"github.com/kaikeventura/cat-runner/src/app/model"
)

type RunnerService struct{}

func ConstructRunnerService() RunnerService {
	return RunnerService{}
}

func RunHttp(httpRunner model.HttpRunner) string {
	url := buildUrl(httpRunner.Http)

	return url
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
