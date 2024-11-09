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

func TestAddGroupTwice(t *testing.T) {
	sut := inv{}
	sut.add_group("foo")
	sut.add_group("foo")

	if len(sut.groups) != 1 {
		t.Fatal("wrong numbers of goups")
	}
}