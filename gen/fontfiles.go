package gen

import (
	"strings"

	"github.com/clark-john/gfontsgen/config"
	"github.com/clark-john/gfontsgen/font"
	"github.com/clark-john/gfontsgen/json"
	"github.com/fatih/color"
	"github.com/samber/lo"
)

func _GenerateFontFiles(_font json.FontItem, options *GenerateFontFileOptions){
	for _, variant := range options.Variants {
		v := ReplaceIWithItalic(variant)
		if HasKey(_font.Files, v) {
			Download(_font.Files[v], options.Path, RemoveSpaces(_font.Family), v)
		} else {
			color.HiYellow("Warning: variant %s doesn't exist on font %s", v, _font.Family)
		}
	}
}

func GenerateFontFiles(_font json.FontItem, options *GenerateFontFileOptions) {
	if IsAll(options.Variants) {
		for _, variant := range lo.Keys(_font.Files) {
			v := ReplaceIWithItalic(variant)
			Download(_font.Files[v], options.Path, _font.Family, v)
		}
	} else {
		_GenerateFontFiles(_font, options)
	}
}

func GenerateFontFilesMultiple(itemsResponse []json.FontItem, options *GenerateFontFilesMultiOptions){
	lo.ForEach(options.OptionItems, func (item config.OptionItem, index int) {
		fontItem, isFound := lo.Find(itemsResponse, func (fitem json.FontItem) bool {
			return strings.EqualFold(fitem.Family, item.FontFamily)
		})
		if isFound {
			_GenerateFontFiles(fontItem, &GenerateFontFileOptions{
				Path: options.Path,
				Variants: item.Variants,
			})
		} else {
			color.HiYellow(`Ignored provided font family "%s" from configuration since it doesn't exist from the list of available fonts.`, item.FontFamily)
		}
	})
}

func Download(url string, path string, family string, variant string){
	font.DownloadFile(font.DownloadOptions{
		Url: url,
		Path: path,
		Family: family,
		Variant: variant,
	})
}

func RemoveSpaces(text string) string {
  return strings.ReplaceAll(text, " ", "")
}
