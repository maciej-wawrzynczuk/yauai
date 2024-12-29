package db2inv

import (
	"yauai/db"
	"yauai/inv"
)

const default_group = "ungroupped"


type Db2inv struct {
	host_key string
	group_keys []string
	i inv.Inv
}

func MkInv(db *db.Db, host_key string, group_keys ...string) (*inv.Inv, error)  {
	di := New(host_key, group_keys...)
	for _, e := range *db {
		err := di.AddEntry(e)
		if err != nil {
			return nil, err
		}
	}
	return di.Inv(), nil
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
	var group0 string
	if len(di.group_keys) > 0 {
		group0 = e[di.group_keys[0]]
	} else {
		group0 = default_group
	}
	
	di.i.AddGroup(group0)
	hostname := e[di.host_key]
	err := di.i.AddHost(hostname, group0)
	return err
}

func (di *Db2inv) Inv() *inv.Inv{
	return &di.i
}
