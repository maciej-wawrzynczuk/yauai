package testutils

import (
	"yauai/inv"
	"encoding/json"
)


type Hosts map[string]interface{}
type Group map[string]Hosts
type Groups map[string]Group
type Meta interface{}

func Extract(i inv.Inv) (*Groups, *Meta, error) {
	text, err := json.Marshal(i)
	if err != nil {
		return nil, nil, err
	}
	any := new(Groups)
	err = json.Unmarshal(text, any)
	if err != nil {
		return nil, nil, err
	}
	return any, nil, nil
}
