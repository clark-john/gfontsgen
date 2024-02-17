package http

import (
	"bytes"
	"io"
)

func ReadBody(r io.Reader) string {
	b, err := io.ReadAll(r)

	if err != nil {
		return ""
	}

	return bytes.NewBuffer(b).String()
}
