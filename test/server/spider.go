package server

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type Spider struct {
	method string
	client *http.Client
	Header map[string]string
	url    string
	param  map[string]string
}

func NewSpider() *Spider {
	return &Spider{
		client: &http.Client{},
		method: "",
		Header: make(map[string]string),
		param:  make(map[string]string),
	}
}

func (s *Spider) Method(method string) *Spider {
	s.method = method
	return s
}

func (s *Spider) Defalut() *Spider {
	s.Header["User-Agent"] = "Mozilla/5.0"
	s.Header["Content-Type"] = "application/json"
	return s
}

func (s *Spider) URL(url string) *Spider {
	s.url = url
	return s
}

func (s *Spider) Param(params map[string]string) *Spider {
	s.param = params
	return s
}

func (s *Spider) Run() *http.Response {
	params, err := json.Marshal(s.param)
	Panic(err)
	req, err := http.NewRequest(s.method, s.url, bytes.NewReader(params))
	for k, v := range s.Header {
		req.Header.Add(k, v)
	}
	Panic(err)
	res, err := s.client.Do(req)
	Panic(err)
	return res
}
