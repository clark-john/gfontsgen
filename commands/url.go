package commands

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/clark-john/gfontsgen/font"
	"github.com/clark-john/gfontsgen/http"
	"github.com/clark-john/gfontsgen/json"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const FONT_URL = "https://fonts.googleapis.com/css"

func SendAndEncode(h *http.HttpClient) (*json.FullResponse, bool) {
  str := h.Send()
  if strings.Contains(str, "404") {
    return nil, false
  }
  resp, _ := json.EncodeResponseStringJson(str)
  return resp, true
}

/**
 * gen's subcommand for generating urls that can be used for css
*/
func UrlCommand() *cobra.Command {
  var variants string
  var copyToClipboard bool
  var cssImport bool
  client := http.NewHttpClientWithKey(key.GetApiKey())

  com := &cobra.Command{
    Use: "url [FontName]",
    Short: "Generate URLs for css usage",
  }

  fl := com.Flags()

  fl.StringVarP(&variants, "variants", "v", "", "Generate with specific variants (e.g. 400, 600, 200i, etc)")
  fl.BoolVar(&copyToClipboard, "copy", false, "Copy URL to clipboard")
  fl.BoolVarP(&cssImport, "css-import", "i", false, "Display as CSS import rule")

  com.Run = func(cmd *cobra.Command, args []string) {
    if len(args) < 1 {
      color.HiYellow("Provide a font family name")
      utils.Exit(1)
    }

    client.SetQuery("family", utils.ToTitleCase(args[0]))

    options := DefaultGenUrlOptions()

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

    resp, isFound := SendAndEncode(client)

    if isFound {
      generated, hasWarnings := GenerateUrl(resp.Items[0], options)

      if cssImport {
        generated = `@import url("` + generated + `");`
      }

      if !hasWarnings {
        fmt.Println(generated)
  
        if copyToClipboard {
          err := clipboard.WriteAll(generated)
          if err == nil {
            color.HiGreen("Successfully copied!")
          } else {
            color.HiRed("Failed to copy")
          }
        }
      }


    } else {
      fmt.Println(color.HiRedString("Font family not found"))
    }
  }

  return com
}

func GenerateUrl(font json.FontItem, options *GenerateUrlOptions) (string, bool) {
  u, _ := url.Parse(FONT_URL)
  var hasWarnings []int

  var family strings.Builder

  family.WriteString(font.Family)  
  family.WriteString(":")

  if !IsAll(options.Variants) {
    for i, variant := range options.Variants {
      family.WriteString(variant)

      /* error */
      v := ReplaceIWithItalic(variant)

      if !HasKey(font.Files, v) {
        color.HiYellow(
          "Warning: variant %s doesn't exist on font %s.", 
          v, font.Family,
        )
        hasWarnings = append(hasWarnings, 1)
      }

      if i < len(options.Variants) - 1 {
        family.WriteString(",")
      }
    }
  } else {
    family.WriteString(strings.Join(lo.Map(font.Variants, func(item string, i int) string {
      if item != "italic" {
        return strings.ReplaceAll(item, "italic", "i")
      } else {
        return "italic"
      }
    }),","))
  }

  q := u.Query()

  q.Set("family", family.String())
  q.Set("display", "swap")

  u.RawQuery = q.Encode()

  return UnescapeColonAndComma(u.String()), len(hasWarnings) > 0
}

func UnescapeColonAndComma(str string) string {
  return strings.ReplaceAll(
    strings.ReplaceAll(str, "%3A", ":"), "%2C", ",",
  )
}

func IsAll(vars []string) bool {
  return len(lo.Filter(vars , func(item string, index int) bool {
    return "all" == strings.ToLower(item)
  })) > 0
}
