package inv

import (
	"encoding/json"
	"errors"
)
type Inv struct {
	groups map[string]Group
	meta meta
}

type Group struct{
	Hosts map[string]vars `json:"hosts"`
}
type meta struct{}
type vars map[string]string

func NewInv() *Inv{
	i := Inv{}
	i.groups = make(map[string]Group)
	return &i
}

func (i Inv) MarshalJSON() ([]byte, error) {
	dirty_inv := make(map[string]interface{})
	for name, group := range i.groups {
		dirty_inv[name] = group
	}
	dirty_inv["_meta"] = i.meta

	result, err := json.Marshal(dirty_inv)
	return result, err
}

func (i Inv) GetGroup(name string) (*Group, error) {
	g, ok := i.groups[name]
	if !ok {
		return nil, errors.New("no such group")
	}
	return &g, nil
}

func (i *Inv) AddGroup(name string) {
	_, ok := i.groups[name]
	if !ok {
		i.groups[name] = *newGroup()
	}
}

func (i *Inv) AddHost(h_name string, g_name string) error {
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

func (g Group) HostNames() []string {
	result := []string{}
	for h := range g.Hosts {
		result = append(result, h)
	}
	return result
}