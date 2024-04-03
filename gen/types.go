package gen

import "github.com/clark-john/gfontsgen/config"

type GenerateUrlOptions struct {
	Copy bool
	Variants []string
}

type GenerateFontFileOptions struct {
	Variants []string
	Path string
}

type GenerateFontFilesMultiOptions struct {
	OptionItems []config.OptionItem
	Path string
}

func DefaultGenUrlOptions() *GenerateUrlOptions {
	return &GenerateUrlOptions{
		Copy: false,
		Variants: []string{"regular"},
	}
}

func NewGenFontFileMultiOptions(path string, ops []config.OptionItem) *GenerateFontFilesMultiOptions {
	return &GenerateFontFilesMultiOptions{
		Path: path,
		OptionItems: ops,
	}
}

func DefaultGenFfOptions() *GenerateFontFileOptions {
	return &GenerateFontFileOptions{
		Variants: []string{"regular"},
		Path: "fonts",
	}
}
