package main

import (
	"testing"
)

func TestHaMeta(t *testing.T) {
	i := newInv()
	_, ok := i["_meta"]
	if !ok {
		t.Fatal("no _meta")
	}
}

func TestAddGroup(t *testing.T) {
	i := newInv()
	i.add_group("foo")

	_, ok := i["foo"]
	if !ok {
		t.Fatal("key not found")
	}
}
