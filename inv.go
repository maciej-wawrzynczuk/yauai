package main

import "encoding/json"

type inv map[string]interface{}

func newInv() inv {
	i := make(inv)
	i.add_meta()
	return i
}

func (i inv) add_group(name string) {
	i[name] = nil
}

func (i inv) mk_json() (string, error) {
	j, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (i inv) add_meta() {
	i["_meta"] = map[string]interface{}{"hostvars": nil}
}
