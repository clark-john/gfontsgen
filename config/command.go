package config

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

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
