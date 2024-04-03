package commands

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

type OptionItem struct {
	FontFamily string
	Variant string
}

type Config struct {
	Options []OptionItem
	Woff bool
	DeleteFontDir bool
	Copy bool
	ToCssImport bool
}

func ConfigFlag(fset *pflag.FlagSet, variable *string){
	fset.StringVar(variable, "config", "", "Url/font gen config file to use")
}

func IsOnlyConfigFlagOrExit(fset *pflag.FlagSet){
	if fset.Changed("config") {
		return
	} else if fset.Changed("config") && fset.NFlag() > 1 {
		color.HiRed("Cannot use different flags other than config flag")
		os.Exit(1)
	}
}
