package main

import (
	"testing"
)

func TestHaMeta(t *testing.T) {
	i := newInv()
	meta, ok := i["_meta"]
	if !ok {
		t.Fatal("no _meta")
	}

	metas, ok := meta.(map[string]interface{})
	if ! ok {
		t.Fatal("oops, assertion failed")
	}

	_, ok = metas["hostvars"]
	if !ok {
		t.Fatal("no hostvars")
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
