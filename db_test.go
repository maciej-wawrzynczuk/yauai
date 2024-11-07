package main

import (
	"testing"
)

var a_data = []byte(`
[{"foo": "bar"}]
`)


func TestHosts(t *testing.T) {
	sut := db{
		{"foo": "bar"},
	}

	result := sut.hosts("foo")

	if result[0] != "bar" {
		t.Fatal("bad value")
	}
}

func TestNewdb(t *testing.T) {
	sut, err := new_db(a_data)
	if err != nil {
		t.Fatal(err)
	}
	if sut[0]["foo"] != "bar" {
		t.Fatal("bad value")
	}
}