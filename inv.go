package main

import "encoding/json"

type inv struct {
	groups map[string]group
	meta meta
}
type group struct{
	hosts map[string]vars
}
type meta struct{}
type vars map[string]string

func (i inv) MarshalJSON() ([]byte, error) {
	dirty_inv := make(map[string]interface{})
	for name, group := range i.groups {
		dirty_inv[name] = group
	}
	dirty_inv["_meta"] = i.meta

	result, err := json.Marshal(dirty_inv)
	return result, err
}