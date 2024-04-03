package config

import (
	"bytes"
	"encoding/json"
)

type OptionItem struct {
	FontFamily string
	Variants []string
}

func (item OptionItem) String() string {
	b, err := json.Marshal(item)
	if err != nil {
		return ""
	}
	return bytes.NewBuffer(b).String()
}

type ConfigInit struct {
	Data []byte
	Path string
}

type Config struct {
	Options []OptionItem
	Woff bool
	DeleteFontDir bool
	Copy bool
	ToCssImport bool
	OutputPath string
}
