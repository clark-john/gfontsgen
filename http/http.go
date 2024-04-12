package http

import "net/http"

const API_URL string = "https://www.googleapis.com/webfonts/v1/webfonts"

type HttpClient struct {
	Url string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{ Url: API_URL }
}

func NewHttpClientWithKey(key string) *HttpClient {
	url := API_URL
	SetUrlQuery(&url, "key", key)
	return &HttpClient{ Url: url }
}

func (h *HttpClient) SetQuery(key string, value string) {
	SetUrlQuery(&h.Url, key, value)
}

func (h *HttpClient) SetApiKey(key string) {
	h.SetQuery("key", key)
}

func (h *HttpClient) Send() string {
	resp, err := http.Get(h.Url)
	
	if err != nil {
		return err.Error()
	}

	return ReadBody(resp.Body)
}
