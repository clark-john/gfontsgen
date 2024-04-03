package commands

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	conf "github.com/clark-john/gfontsgen/config"
	"github.com/clark-john/gfontsgen/gen"
	"github.com/clark-john/gfontsgen/http"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

/**
 * gen's subcommand for generating font files
 */
func FontFilesCommand() *cobra.Command {
	var variants string
	var outputPath string
	var config string
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
	ConfigFlag(fset, &config)

	com.Run = func(cmd *cobra.Command, args []string) {
    isConfigFlag := conf.IsPresent(cmd)

    IsOnlyConfigFlagOrExit(fset)

		if !isConfigFlag {
			if len(args) < 1 {
	      color.HiYellow("Provide a font family name")
	      utils.Exit(1)
	    }
	    FontfilesGenerateSingle(args[0], client, isWoff, cmd, variants, outputPath)
		} else {
			FontfilesGenerateMultiple(config, client)
		}
	}

	return com
}

func FontfilesGenerateSingle(
	family string,
	client *http.HttpClient,
	isWoff bool,
	cmd *cobra.Command,
	variants string,
	outputPath string,
){
	fontName := utils.ToTitleCase(family)

  client.SetQuery("family", fontName)

  if isWoff {
  	client.SetQuery("capability", "woff2")
  }

  options := gen.DefaultGenFfOptions()

  CheckVariants(
  	cmd.Flag("variants").Changed, variants, &options.Variants,
  )

  resp, isFound := SendAndEncode(client)

  if cmd.Flag("path").Changed {
  	options.Path = outputPath
  }

  if isFound {
    gen.GenerateFontFiles(resp.Items[0], options)
  } else {
    fmt.Println(color.HiRedString("Font family not found"))
  }
}

func FontfilesGenerateMultiple(config string, client *http.HttpClient) {
	cfg := conf.ParseAndValidateConfig(config)
	path := func() string {
		p := cfg.OutputPath
		if p == "" {
			return "fonts"
		}
		return p
	}()
  ops := cfg.Options

  /* reused from url.go line 146 */
  for index, op := range ops {
    if index == len(ops) - 1 {
      CheckVariantsP(op.Variants, true, op.FontFamily)
    } else {
      CheckVariantsP(op.Variants, false, op.FontFamily)
    }
  }

  if cfg.DeleteFontDir {
  	DeleteDirIfExist(path)
  }

  res, _ := SendAndEncode(client)
  /* end of reused block of code */

  gen.GenerateFontFilesMultiple(
  	res.Items, 
  	gen.NewGenFontFileMultiOptions(path, ops),
  )
}

func DeleteDirIfExist(path string){
	_, err := os.Open(path)

	if err == nil {
		if filepath.Walk(path, func(_path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				os.RemoveAll(_path)
			}
			return err
		}) != nil {
			return
		}
	}
}
