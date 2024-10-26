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

func (i inv) mk_json() (string, error) {
	j, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func main() {
	i := mk_inv()
	j, err := i.mk_json()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(j)
}

