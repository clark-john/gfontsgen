package http

import "net/http"

const API_URL string = "https://www.googleapis.com/webfonts/v1/webfonts"

type HttpClient struct {
	url string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{ url: API_URL }
}

func NewHttpClientWithKey(key string) *HttpClient {
	url := API_URL
	SetUrlQuery(&url, "key", key)
	return &HttpClient{ url: url }
}

func (h *HttpClient) SetQuery(key string, value string) {
	SetUrlQuery(&h.url, key, value)
}

func (h *HttpClient) SetApiKey(key string) {
	h.SetQuery("key", key)
}

func (h *HttpClient) Send() string {
	resp, err := http.Get(h.url)
	
	if err != nil {
		return err.Error()
	}

	return ReadBody(resp.Body)
}
