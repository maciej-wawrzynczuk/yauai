package main

import (
	"testing"
)

func TestAddGroup(t *testing.T) {
	sut := newInv()
	sut.add_group("foo")

	_, err := sut.get_group("foo")

	if err != nil {
		t.Fatal("add group don't add group")
	}
}

// TODO: Actually I want to check if adding group twice flattens the existing one.
func TestAddGroupTwice(t *testing.T) {
	sut := newInv()
	sut.add_group("foo")
	sut.add_group("foo")

	if len(sut.groups) != 1 {
		t.Fatal("wrong numbers of goups")
	}
}

func TestAddHostNoSuchGroup(t *testing.T) {
	sut := newInv()
	err := sut.add_host("h", "g")
	if err == nil {
		t.Fatal("it should be error")
	}
}

func TestReallyAddHost(t *testing.T) {
	sut := newInv()
	group_name := "a group"
	host_name := "a host"
	sut.add_group(group_name)
	err := sut.add_host(host_name, group_name)
	if err != nil {
		t.Fatal("there should be no error this time")
	}
	g, ok := sut.groups[group_name]
	if !ok {
		t.Fatal("no group")
	}
	_, ok = g.hosts[host_name]
	if !ok {
		t.Fatal("no host in group")
	}
}