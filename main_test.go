package main

import(
	"testing"
)

func TestAddGroup(t *testing.T) {
	i := mk_inv()
	i.add_group("foo")

	_, ok := i["foo"]
	if ! ok {
		t.Fatal("key not found")
	}
}