package json

import "encoding/json"

func EncodeResponseStringJson(stringJson string) (*FullResponse, error) {
	var resp FullResponse

	err := json.Unmarshal([]byte(stringJson), &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
