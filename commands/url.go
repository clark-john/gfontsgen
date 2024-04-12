package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	conf "github.com/clark-john/gfontsgen/config"
	"github.com/clark-john/gfontsgen/gen"
	"github.com/clark-john/gfontsgen/http"
	"github.com/clark-john/gfontsgen/json"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type GenerateSingleArgs struct {
  Family string
  Client *http.HttpClient
  Cmd *cobra.Command
  Variants string
  CssImport bool
  CopyToClipboard bool
}

type GenerateMultipleArgs struct {
  Config string
  Client *http.HttpClient
}

const FONT_URL = "https://fonts.googleapis.com/css"

func SendAndEncode(h *http.HttpClient) (*json.FullResponse, bool) {
  str := h.Send()
  if strings.Contains(str, "no such host") {
    fmt.Println("Cannot fetch. Check your internet connection.")
    os.Exit(1)
  }  
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
  var config string
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
  conf.ConfigFlag(fl, &config)

  com.Run = func(cmd *cobra.Command, args []string) { 
    isConfigFlag := conf.IsPresent(cmd)

    conf.IsOnlyConfigFlagOrExit(fl)

    if !isConfigFlag {
      if len(args) < 1 {
        color.HiYellow("Provide a font family name")
        utils.Exit(1)
      }
      UrlGenerateSingle(GenerateSingleArgs{
        Cmd: cmd,
        Client: client,
        Family: args[0],
        Variants: variants,
        CssImport: cssImport,
        CopyToClipboard: copyToClipboard,
      })
    } else {
      UrlGenerateMultiple(GenerateMultipleArgs{
        Config: config,
        Client: client,
      })
    }
  }

  return com
}

func UrlGenerateSingle(args GenerateSingleArgs){
  options := gen.DefaultGenUrlOptions()

  family := args.Family
  client := args.Client
  cmd := args.Cmd
  variants := args.Variants
  cssImport := args.CssImport
  copyToClipboard := args.CopyToClipboard

  /* setting the family query to filter it to single font family */
  client.SetQuery("family", utils.ToTitleCase(family))
 
  CheckVariants(cmd.Flag("variants").Changed, variants, &options.Variants)

  resp, isFound := SendAndEncode(client)

  if isFound {
    generated, hasWarnings := gen.GenerateUrl(resp.Items[0], options)

    if cssImport {
      CssImportSurround(&generated)
    }

    if !hasWarnings {
      fmt.Println(generated)

      if copyToClipboard {
        CopyText(generated)
      }
    }

  } else {
    /* this condition is only applied to single family (without config) */
    fmt.Println(color.HiRedString("Font family not found"))
  }
}

func UrlGenerateMultiple(args GenerateMultipleArgs){
  config := args.Config
  client := args.Client

  cfg := conf.ParseAndValidateConfig(config)
  ops := cfg.Options
  
  /* check variants of every option item */
  for index, op := range ops {
    if index == len(ops) - 1 {
      CheckVariantsP(op.Variants, true, op.FontFamily)
    } else {
      CheckVariantsP(op.Variants, false, op.FontFamily)
    }
  }

  res, _ := SendAndEncode(client)
  items := res.Items
  generated, hasWarnings := gen.GenerateUrlMultiple(items, cfg.Options)

  if !hasWarnings {
    if cfg.ToCssImport {
      CssImportSurround(&generated)
    }

    fmt.Println(generated)

    if cfg.Copy {
      CopyText(generated)
    }
  }
}

func CssImportSurround(generated *string) {
  *generated = fmt.Sprintf(`@import url("%s");`, *generated)
}

func CopyText(text string){
  err := clipboard.WriteAll(text)
  if err == nil {
    color.HiGreen("Successfully copied!")
  } else {
    color.HiRed("Failed to copy")
  }
}
