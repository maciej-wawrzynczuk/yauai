package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"yauai/db"
	"yauai/inv"
)

func main() {
	// TODO determin hostname key
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db, err := db.NewDb(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i, err := Db2inv(*db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	text, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} 
	fmt.Print(string(text))
}


func Db2inv(_ db.Db) (*inv.Inv, error) {
	return inv.NewInv(), nil
} 