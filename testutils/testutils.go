package testutils

import (
	"yauai/inv"
	"encoding/json"
)


type Hosts map[string]interface{}
type Group map[string]Hosts
type Inv2 map[string]Group

func Extract(i inv.Inv) (*Inv2, error) {
	text, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	any := new(Inv2)
	err = json.Unmarshal(text, any)
	if err != nil {
		return nil, err
	}
	return any, nil
}
