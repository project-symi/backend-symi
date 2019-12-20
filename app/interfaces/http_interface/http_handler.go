package http_interface

type HttpHandler interface {
	NewHttpRequest(string, string, []byte) error
	SetHeader(string, string)
	DoRequest() error
}

type Request interface {
	Header
}

type Header interface {
	Set(string, string)
}

type Response interface {
	Body
}

type Body interface {
	Close()
}
