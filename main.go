package main

import (
	"fmt"
	"log"
)


func main() {
	i := newInv()
	i.add_group("foo")
	j, err := i.mk_json()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(j)
}
