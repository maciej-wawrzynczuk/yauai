package main

import (
	"encoding/json"
	"fmt"
	"log"
)


func main() {
	i := inv{}

	text, err := json.Marshal(i)

	if err != nil {
		log.Fatal(err)
	} 
	fmt.Print(string(text))
}
