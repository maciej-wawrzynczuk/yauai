package inv

import (
	"encoding/json"
	"errors"
)
type inv struct {
	groups map[string]Group
	meta meta
}

type Group struct{
	// TODO: try to get rid of public members
	Hosts map[string]vars
}
type meta struct{}
type vars map[string]string

func NewInv() *inv{
	i := inv{}
	i.groups = make(map[string]Group)
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

func (i inv) GetGroup(name string) (*Group, error) {
	g, ok := i.groups[name]
	if !ok {
		return nil, errors.New("no such group")
	}
	return &g, nil
}

func (i *inv) AddGroup(name string) {
	i.groups[name] = *newGroup()
}

func (i *inv) AddHost(h_name string, g_name string) error {
	group, ok := i.groups[g_name]
	if !ok {
		return errors.New("no group")
	}
	group.Hosts[h_name] = make(vars)
	return nil
}

func newGroup() *Group {
	newgrp := Group{}
	newgrp.Hosts = make(map[string]vars)
	return &newgrp	
}