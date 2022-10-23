package bejson

import (
	"encoding/json"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//UnmarshalString parse json string to output struct
func UnmarshalString(jsonString string, output interface{}) error {
	err := json.Unmarshal([]byte(jsonString), output)
	if err != nil {
		return err
	}
	return nil
}

//Marshal parse struct to bejson.JSON
func Marshal(input interface{}) (*JSON, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, grerrors.NewDefaultError(err)
	}

	output := JSON(b)
	return &output, nil
}