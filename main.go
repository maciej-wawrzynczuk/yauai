package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type inv map[string]interface{}

func mk_inv() inv {
	i := make(inv)
	return i
}

func main() {
	i := mk_inv()
	j, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}

