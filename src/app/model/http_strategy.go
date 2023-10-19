package model

type Strategy struct {
	Http        Http
	VirtualUser VirtualUser
}

type Http struct {
	Protocol   Protocol
	Host       string
	Port       int
	HttpMethod HttpMethod
	Timeout    int
	Headers    []KVParam
	Parameters []KVParam
	Body       Body
}

type Protocol string

const (
	HTTP  Protocol = "GET"
	HTTPS Protocol = "POST"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PATCH  HttpMethod = "PATCH"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

type KVParam struct {
	Key   string
	Value string
}

type Body struct {
	BodyFormat  BodyFormat
	ContentText string
}

type BodyFormat string

const (
	JSON BodyFormat = "JSON"
)

type VirtualUser struct {
	UsersAmount        int
	InteractionsAmount int
	InteractionDelay   int
}
