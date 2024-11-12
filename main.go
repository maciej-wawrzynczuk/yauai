package main

import (
	"encoding/json"
	"fmt"
	"log"
	"yauai/inv"
)

// TODO: Move stuff to packages and use public interface only

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
