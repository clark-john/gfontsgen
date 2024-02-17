package main

import (
	"fmt"
	"os"

	"github.com/clark-john/gfontsgen/commands"
	"github.com/clark-john/gfontsgen/key"
	"github.com/clark-john/gfontsgen/utils"
	"github.com/spf13/cobra"
)

const KEY_NAME string = key.KEY_NAME
const Version string = "1.0.0"

func main() {
	if !utils.IsInEnviron(KEY_NAME) {
		fmt.Println("Set your environment variable " + KEY_NAME + " to your API key you obtained from the Google Fonts API")
		fmt.Println("Get your API key here: https://developers.google.com/fonts/docs/developer_api#APIKey")
		return
	}

	com := cobra.Command{
		Use: "gfontsgen",
		Version: Version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	com.AddCommand(
		commands.ListCommand(),
		commands.GenCommand(),
	)

	err := com.Execute()

	if err != nil {
		os.Exit(1)
	}
}

/*
commands:

gfontsgen gen fontfiles <FamilyName> # generate ttf (default)
gfontsgen gen fontfiles --woff # generate woff (false by default)
gfontsgen gen fontfiles -p myfonts # or --path # generate files to fonts folder (default: fonts)
gfontsgen gen fontfiles -v 200i,400 # same lng sa kanina

*/