package model

type HttpRunner struct {
	RequestName string
	Http        Http
	VirtualUser VirtualUser
}

type Http struct {
	Protocol   Protocol
	Host       string
	Port       int
	Path       string
	HttpMethod HttpMethod
	Timeout    int
	Headers    []KVParam
	Parameters []KVParam
	Body       Body
}

type Protocol string

const (
	HTTP  Protocol = "HTTP"
	HTTPS Protocol = "HTTPS"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
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
