package main

import (
	"encoding/json"
	"fmt"
	"log"
)

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
	i["_meta"] = nil
}

func main() {
	i := newInv()
	i.add_group("foo")
	j, err := i.mk_json()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(j)
}