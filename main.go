package main

import (
	"flag"
	"log"
)


func main() {
	var hostname_field string
	flag.StringVar(&hostname_field, "host_field", "", "Field name to be used as hostname")
	flag.Parse()
	log.Println(hostname_field)
}

