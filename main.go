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
	host_key := flag.String("host_key", "", "Key to become hostname")
	input_file := flag.String("input_file", "", "Input file. Stdin if omit")
	flag.Parse()

	if *host_key == "" {
		fmt.Println("Need a hostname field")
		os.Exit(1)
	}

	var f *os.File
	if *input_file != "" {
		var err error
		f, err = os.Open(*input_file)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		f = os.Stdin
	}

	data, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db, err := db.FromJson(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i, err := db2inv.Db2inv_to_drop(*db, *host_key)
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
