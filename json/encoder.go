package json

import (
	"encoding/json"
	"github.com/clark-john/gfontsgen/utils"
)

func EncodeResponseStringJson(stringJson string) (*FullResponse, error) {
	var resp FullResponse

	err := json.Unmarshal(utils.StringToBytes(stringJson), &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
