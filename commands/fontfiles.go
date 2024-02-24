package commands

import (
	"fmt"
	"strings"

	"github.com/clark-john/gfontsgen/font"
	"github.com/clark-john/gfontsgen/http"
	"github.com/clark-john/gfontsgen/json"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

/**
 * gen's subcommand for generating font files
 */
func FontFilesCommand() *cobra.Command {
	var variants string
	var outputPath string
	// var config string
	var isWoff bool

	client := http.NewHttpClientWithKey(key.GetApiKey())

	com := &cobra.Command{
		Use: "fontfiles [FontName]",
		Short: "Generate font files",
	}

	fset := com.Flags()

	fset.StringVarP(&outputPath, "path", "p", "fonts", "Output path to generate font files")
	fset.StringVarP(&variants, "variants", "v", "", "Generate with specific variants (e.g. 400, 600, 200i, etc)")
	fset.BoolVar(&isWoff, "woff", false, "Use WOFF2 format")
	// ConfigFlag(fset, &config)

	com.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
      color.HiYellow("Provide a font family name")
      utils.Exit(1)
    }

    fontName := utils.ToTitleCase(args[0])

    client.SetQuery("family", fontName)

    if isWoff {
    	client.SetQuery("capability", "woff2")
    }

    options := DefaultGenFfOptions()

    /* reused from url.go */
    if cmd.Flag("variants").Changed {
      if !font.ValidateVariantsArg(variants) {
        color.HiRed(`Invalid variants argument it must be in this format: "400,500,600"`)
        utils.Exit(1)
      }
      v := strings.Split(variants, ",")
      
      indices := font.ValidateVariants(v)

      if indices != nil {
        font.PrintVariantsErrAndExit(v, indices)
      }

      options.Variants = v
    }
    /* end fo reused block */

    resp, isFound := SendAndEncode(client)

    if isFound {
	    GenerateFontFiles(resp.Items[0], options)
    } else {
      fmt.Println(color.HiRedString("Font family not found"))
    }
	}

	return com
}

func GenerateFontFiles(_font json.FontItem, options *GenerateFontFileOptions) {
	if IsAll(options.Variants) {
		for _, variant := range lo.Keys(_font.Files) {
			v := ReplaceIWithItalic(variant)
			Download(_font.Files[v], options.Path, _font.Family, v)
		}
	} else {
		for _, variant := range options.Variants {
			v := ReplaceIWithItalic(variant)
			if HasKey(_font.Files, v) {
				Download(_font.Files[v], options.Path, _font.Family, v)
			} else {
				color.HiYellow("Warning: variant %s doesn't exist on font %s", v, _font.Family)
			}
		}
	}
}

func Download(url string, path string, family string, variant string){
	font.DownloadFile(font.DownloadOptions{
		Url: url,
		Path: path,
		Family: family,
		Variant: variant,
	})
}

func ReplaceIWithItalic(str string) string {
	bef, _, f := strings.Cut(str, "i")
	if !f {
		return bef
	} else {
		return bef + "italic"
	}
}

func HasKey(m map[string]string, value string) bool {
	for k := range m {
		if k == value {
			return true
		}
	}
	return false
}
