package commands

import "github.com/spf13/cobra"

type GenerateUrlOptions struct {
	Copy bool
	Variants []string
}

type GenerateFontFileOptions struct {
	Variants []string
	Path string
}

func DefaultGenUrlOptions() *GenerateUrlOptions {
	return &GenerateUrlOptions{
		Copy: false,
		Variants: []string{"regular"},
	}
}

func DefaultGenFfOptions() *GenerateFontFileOptions {
	return &GenerateFontFileOptions{
		Variants: []string{"regular"},
		Path: "fonts",
	}
}

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
