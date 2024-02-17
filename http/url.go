package http

import "net/url"

func SetUrlQuery(_url *string, key string, value string) {
	u, err := url.Parse(*_url)
	
	if err != nil {
		return
	}

	q := u.Query()
	q.Set(key, value)

	u.RawQuery = q.Encode()

	*_url = u.String()
}
