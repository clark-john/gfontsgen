package key

import (
	"os"
	"github.com/clark-john/gfontsgen/utils"
)

const KEY_NAME string = "GFONTSGEN_API_KEY"

/**
 *  only used if GFONTSGEN_API_KEY is actually set 
*/
func GetApiKey() string {
	m := utils.ToMap(os.Environ())
	return m[KEY_NAME]
}
