package db2inv

import(
	"yauai/db"
	"yauai/inv"
)

func Db2inv(db db.Db, host_key string, groups ...string) (*inv.Inv, error) {
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