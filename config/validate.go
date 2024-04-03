package config

import (
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

const (
	MISSING_REQUIRED = "Missing required property: "
	VARIANTS_MINIMUM = "Variants must have at least one item"
)

func PrintErrorMapAndExit(mp map[int][]string){
	fmt.Println("Following option items are invalid:")
	for key, val := range mp {
		fmt.Printf("  At index %d:\n", key)
		for _, err := range val {
			fmt.Printf("    %s\n", err)
		}
	}
	os.Exit(1)
}

func JsonPathOptionItem(index int, property string) string {
	return fmt.Sprintf("options.%d.%s", index, property)
}

func CheckAllOptionItems(items []OptionItem, config *ConfigInit) {
	errorMap := make(map[int][]string)
	cfg := config.Data
	for index, item := range items {
		ff := gjson.GetBytes(
			cfg, JsonPathOptionItem(index, "fontFamily"),
		).Exists()
		varnts := gjson.GetBytes(
			cfg, JsonPathOptionItem(index, "variants"),
		).Exists()

		if !varnts && !ff {
			errorMap[index] = []string{
				MISSING_REQUIRED + "fontFamily",
				MISSING_REQUIRED + "variants",
			}
		} else if !ff {
			errorMap[index] = []string{
				MISSING_REQUIRED + "fontFamily",
			}			
		} else if !varnts {
			errorMap[index] = append(errorMap[index], MISSING_REQUIRED + "variants")
		}

		if varnts {
			if len(item.Variants) == 0 {
				errorMap[index] = append(errorMap[index], VARIANTS_MINIMUM)
			}
		}
	}
	if len(errorMap) > 0 {
		PrintErrorMapAndExit(errorMap)
	}
}

/**
 * This only validates if required options are satisifed (i.e. values aren't validated)
 */
func ValidateConfig(config *ConfigInit) *Config {
	cfg := config.GetDataString()

	if !gjson.Get(cfg, "options").Exists() {
		fmt.Println("Missing required property: options")
		os.Exit(1)
	}

	conf := config.ToConfig()

	if len(conf.Options) == 0 {		
		fmt.Println("Options must have at least one item.")
		os.Exit(1)
	}

	CheckAllOptionItems(conf.Options, config)

	return conf
}

