package db_test

import (
	"testing"
	"yauai/db"
)

var a_data = []byte(`
[{"foo": "bar"},
 {"baz": "qux"}]
`)

func TestHosts(t *testing.T) {
	var in = []byte(`[{"foo":"bar"}]`)
	sut, err := db.FromJson(in)
	if err != nil {
		t.Fatal(err)
	}

	result := sut.Hosts("foo")

	if result[0] != "bar" {
		t.Fatal("bad value")
	}
}

func TestNewdb(t *testing.T) {
	sut, err := db.FromJson(a_data)
	if err != nil {
		t.Fatal(err)
	}
	if len(*sut) != 2 {
		t.Fatal("bad size")
	}
	if (*sut)[0]["foo"] != "bar" {
		t.Fatal("bad value")
	}
}
