package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/clark-john/gfontsgen/http"
	"github.com/clark-john/gfontsgen/json"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func ListCommand() *cobra.Command {
	key := key.GetApiKey()
	client := http.NewHttpClientWithKey(key)

	var classification string
	var decorativeStroke string
	var search string
	var writeToFile string
	var limit int

	com := &cobra.Command {
		Use: "list",
		Short: "Get a list of fonts from Google Fonts",
	}

	fset := com.Flags()

	fset.StringVarP(
		&classification, "classification", "c", "", 
		"Classification (display, handwriting, mono/monospace)",
	)
	fset.StringVarP(
		&decorativeStroke, "decorative-stroke", "d", "", 
		"Decorative stroke (serif, sans-serif, slab-serif)",
	)
	fset.StringVarP(
		&search, "search", "s", "", 
		"Search for a font",
	)
	fset.StringVar(
		&writeToFile, "write-to-file", "", 
		"Write output to file",
	)
	fset.IntVarP(
		&limit, "limit", "l", 0, 
		"Limit the number of items",
	)

	writeTfFlag := com.Flag("write-to-file")
	writeTfFlag.NoOptDefVal = "fonts.txt"

	com.Run = func(cmd *cobra.Command, args []string) {
		str := client.Send()

		res, err := json.EncodeResponseStringJson(str)

		if err != nil {
			fmt.Println("An error occurred")
			return
		}

		it := res.Items

		if cmd.Flag("decorative-stroke").Changed {
			if !ValidateDecorativeStroke(decorativeStroke) {
				fmt.Println("Invalid decorative stroke it must be the following: serif, sans-serif")
				return
			}

			it = lo.Filter(it, func(val json.FontItem, i int) bool {
				return decorativeStroke == val.Category
			})
		}

		if cmd.Flag("classification").Changed {
			c := strings.ToLower(classification)

			if c == "mono" {
				c = "monospace";
			}

			if !ValidateClassification(c) {
				fmt.Println("Invalid classification it must be the following: display, handwriting, mono/monospace")
				return
			}

			it = lo.Filter(it, func(val json.FontItem, i int) bool {
				return c == val.Category
			})
		}
		
		if cmd.Flag("search").Changed {
			it = lo.Filter(it, func(val json.FontItem, i int) bool {
				return strings.Contains(
					strings.ToLower(val.Family), strings.ToLower(search),
				)
			})
		}

		if cmd.Flag("limit").Changed {
			it = utils.Limit(it, limit)
		}

		var b strings.Builder

		if len(it) > 0 {
			for index, item := range it {
				b.WriteString("#" + fmt.Sprintf("%d", index + 1) + "\n")
				b.WriteString("Name: " + item.Family + "\n")
				b.WriteString("Variants: " + strings.Join(item.Variants, ", ") + "\n")
				b.WriteString("PreviewURL: " + "https://fonts.google.com/specimen/" + 
					strings.ReplaceAll(item.Family, " ", "+") + "\n\n",
				)
			}
			if !writeTfFlag.Changed {
				fmt.Print(b.String())
			}
		} else {
			fmt.Println("No results")
		}

		if writeTfFlag.Changed {
			err := os.WriteFile(writeToFile, utils.StringToBytes(b.String()), 0666)
			if err != nil {
				fmt.Println("An error occurred")
				os.Exit(1)
			}
		}
	}

	return com
}

func ValidateClassification(cl string) bool {
	validValues := []string{"display", "handwriting", "mono", "monospace"}
	return utils.IsIn(validValues, strings.ToLower(cl))
}

func ValidateDecorativeStroke(stroke string) bool {
	validValues := []string{"serif", "sans-serif"}
	return utils.IsIn(validValues, strings.ToLower(stroke))
}
