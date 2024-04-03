package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	p "path"

	"github.com/spf13/cobra"
)

func (c ConfigInit) ToConfig() *Config {
	var config Config
	err := json.Unmarshal(c.Data, &config)
	if err != nil {
		return nil
	}
	return &config
}

func (c ConfigInit) GetDataString() string {
	return bytes.NewBuffer(c.Data).String()
}

func IsPresent(cmd *cobra.Command) bool {
	return cmd.Flag("config").Changed
}

func GetVariants(option *OptionItem) []string {
	variants := make([]string, len(option.Variants))
	for _, variant := range option.Variants {
		variants = append(variants, variant)
	}
	return variants
}

func OpenConfig(path string) *ConfigInit {
	b, err := os.ReadFile(path)
	if err != nil {
		ExitOne("Cannot find/open config file")
	}
	return &ConfigInit{
		Data: b,
		Path: path,
	}
}

func ParseAndValidateConfig(path string) *Config {
	cinit := OpenConfig(path)

	if p.Ext(cinit.Path) != ".json" {
		ExitOne("File should be in a JSON format")
	}

	if !json.Valid(cinit.Data) {
		ExitOne("Invalid JSON syntax")
	}

	return ValidateConfig(cinit)
}

func ExitOne(message string){
	fmt.Println(message)
	os.Exit(1)
}
