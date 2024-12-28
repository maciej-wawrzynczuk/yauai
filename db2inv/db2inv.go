package db2inv

import (
	"yauai/db"
	"yauai/inv"
)

const default_group = "ungroupped"

// TODO: Make a class.

type Db2inv struct {
	host_key string
	group_keys []string
	i inv.Inv
}

func New(host_key string, group_keys ...string) *Db2inv {
	di := Db2inv{
		host_key: host_key,
		group_keys: group_keys,
		i: *inv.NewInv(),
	}
	return &di
}

func (di *Db2inv) AddEntry(e db.Entry) error{
	di.i.AddGroup(default_group)
	hostname := e[di.host_key]
	err := di.i.AddHost(hostname, default_group)
	return err
}

func (di *Db2inv) Inv() *inv.Inv{
	return &di.i
}
