package db2inv

import (
	"yauai/db"
	"yauai/inv"
)

// TODO: Make a class.

type Db2inv struct {
	host_key string
	group_keys []string
}


func Db2inv_to_drop(db db.Db, host_key string, group_keys ...string) (*inv.Inv, error) {
	default_group := "ungroupped"
	i := inv.NewInv()
	i.AddGroup(default_group)
	for _, name := range db.Hosts(host_key) {
		err := i.AddHost(name, default_group)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}
