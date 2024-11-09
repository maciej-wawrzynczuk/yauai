package main

import (
	"encoding/json"
	"fmt"
	"log"
)


func main() {
	i := inv{}
	i.add_group("foo")
	text, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	} 
	fmt.Print(string(text))
}
