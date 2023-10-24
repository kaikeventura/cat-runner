package model

import "net/http"

type Response struct {
	ResponseStatus      ResponseStatus
	ResponseBody        string
	Time                int
	VirtualUserResponse VirtualUserResponse
	RequestData         RequestData
}

type ResponseStatus struct {
	Description string
	Code        int
}

type VirtualUserResponse struct {
	UserId    int
	RequestId int
}

type RequestData struct {
	Method string
	Url    string
	Header http.Header
	Body   string
}
