package main

import (
	"testing"
)

func TestAddGroup(t *testing.T) {
	sut := inv{}
	sut.add_group("foo")

	_, ok := sut.groups["foo"]

	if !ok {
		t.Fatal("add group don't add group")
	}
}