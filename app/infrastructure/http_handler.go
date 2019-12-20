package infrastructure

import (
	"bytes"
	"net/http"
	"project-symi-backend/app/interfaces/http_interface"
)

type HttpHandler struct {
	request *http.Request
}

func NewHttpHandler() (httpHandler http_interface.HttpHandler) {
	httpHandler = new(HttpHandler)
	return
}

func (handler *HttpHandler) NewHttpRequest(method string, url string, body []byte) (err error) {
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(body),
	)
	handler.request = req
	return
}

func (handler *HttpHandler) SetHeader(name string, headerType string) {
	handler.request.Header.Set(name, headerType)
}

func (handler *HttpHandler) DoRequest() (err error) {
	client := http.Client{}
	_, err = client.Do(handler.request)
	return
}
