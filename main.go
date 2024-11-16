package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"yauai/db"
	"yauai/db2inv"
)

func main() {
	// TODO determin hostname key
	host_key := flag.String("host_key", "", "KEy to become hostname")
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
	i, err := db2inv.Db2inv(*db, *host_key)
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
