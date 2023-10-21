package model

import "net/http"

type Response struct {
	ResponseStatus ResponseStatus
	ResponseBody   string
	Time           int
	RequestData    RequestData
}

type ResponseStatus struct {
	Description string
	Code        int
}

type RequestData struct {
	Method string
	Url    string
	Header http.Header
	Body   string
}
