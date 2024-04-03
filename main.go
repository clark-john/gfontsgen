package main

import (
	"fmt"
	"os"
	"github.com/clark-john/gfontsgen/commands"
	"github.com/clark-john/gfontsgen/key"
	"github.com/spf13/cobra"
)

const KEY_NAME string = key.KEY_NAME
const Version string = "1.1.0"

func main() {
	_, isFound := os.LookupEnv(KEY_NAME)

	if !isFound {
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
		commands.GfontswebCommand(),
	)

	err := com.Execute()

	if err != nil {
		os.Exit(1)
	}
}
