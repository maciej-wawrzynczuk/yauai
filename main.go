package main

import (
	"encoding/json"
	"fmt"
	"log"
	"yauai/inv"
)

func main() {
	i := inv.NewInv()
	i.AddGroup("foo")
	i.AddHost("bar_host", "foo")
	text, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Print(string(text))
}
