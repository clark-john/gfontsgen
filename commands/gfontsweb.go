package commands

import (
	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

const GOOGLE_FONTS_URL = "https://fonts.google.com"

func GfontswebCommand() *cobra.Command {
	return &cobra.Command{
		Use: "gfontsweb",
		Short: "Open Google Fonts website",
		Run: func(cmd *cobra.Command, args []string) {
			err := browser.OpenURL(GOOGLE_FONTS_URL)
			if err == nil {
				color.HiGreen("Opened")
			}
		},
	}
}
