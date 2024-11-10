package main

import (
	"encoding/json"
	"errors"
)
// TODO: Create functions for initialisation
type inv struct {
	groups map[string]group
	meta meta
}
type group struct{
	hosts map[string]vars
}
type meta struct{}
type vars map[string]string

func newInv() *inv{
	i := inv{}
	return &i
}

func (i inv) MarshalJSON() ([]byte, error) {
	dirty_inv := make(map[string]interface{})
	for name, group := range i.groups {
		dirty_inv[name] = group
	}
	dirty_inv["_meta"] = i.meta

	result, err := json.Marshal(dirty_inv)
	return result, err
}

func (i *inv) add_group(g string) {
	if i.groups == nil {
		i.groups = make(map[string]group)
	}
	newgrp := group{}
	newgrp.hosts = make(map[string]vars)
	i.groups[g] = newgrp
}

func (i *inv) add_host(h_name string, g_name string) error {
	group, ok := i.groups[g_name]
	if !ok {
		return errors.New("no group")
	}
	group.hosts[h_name] = make(vars)
	return nil
}