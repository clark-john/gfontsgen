package commands

import (
	"os"
	"strings"

	"github.com/clark-john/gfontsgen/font"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func GenCommand() *cobra.Command {
	com := &cobra.Command{
		Use: "gen",
		Short: "Generate font URL/font files",
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flags().NFlag() < 1 {
				cmd.Help()
			}
		},
	}

	com.AddCommand(
		FontFilesCommand(),
		UrlCommand(),
	)

	return com
}

func CheckVariants(isVariantsFlagPresent bool, varnts string, variants *[]string){
	if isVariantsFlagPresent {
		if !font.ValidateVariantsArg(varnts) {
		  color.HiRed(`Invalid variants argument it must be in this format: "400,500,600"`)
		  os.Exit(1)
		}
		v := strings.Split(varnts, ",")

		indices := font.ValidateVariants(v)

		if indices != nil {
		  font.PrintVariantsErr(v, indices, true)
		}

		*variants = v		
	}
}

func CheckVariantsP(v []string, isExit bool, family string){
	indices := font.ValidateVariants(v)

	if indices != nil {
	  font.PrintVariantsErrWithName(v, indices, isExit, family)
	}
}
