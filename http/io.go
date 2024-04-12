package http

import "io"

func ReadBody(r io.Reader) string {
	b, err := io.ReadAll(r)

	if err != nil {
		return ""
	}

	return string(b)
}
