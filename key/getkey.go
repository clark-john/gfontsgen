package key

import "os"

const KEY_NAME string = "GFONTSGEN_API_KEY"

/**
 *  get currently set api key 
*/
func GetApiKey() string {
	apiKey, _ := os.LookupEnv(KEY_NAME)
	return apiKey
}
