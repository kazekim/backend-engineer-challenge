package bejson

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSON []byte

// String return json in string format
func (j JSON) String() string {
	return string(j)
}

// Value return json in value format
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

// Scan parse value to json format
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON return json in []byte format
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON parse byte data into json
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// Unmarshal parse json into output struct
func (j *JSON) Unmarshal(output interface{}) error {
	err := json.Unmarshal(*j, output)
	if err != nil {
		return err
	}
	return nil
}

// IsNull check if json is null
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals check if value of json is equal to j1
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
