package inv_test

import (
	"testing"
	"yauai/inv"
)

func TestAddGroup(t *testing.T) {
	sut := inv.NewInv()
	sut.AddGroup("foo")

	_, err := sut.GetGroup("foo")

	if err != nil {
		t.Fatal("add group don't add group")
	}
}

// TODO: Actually I want to check if adding group twice flattens the existing one.
func TestAddHostNoSuchGroup(t *testing.T) {
	sut := inv.NewInv()
	err := sut.AddHost("h", "g")
	if err == nil {
		t.Fatal("it should be error")
	}
}

func TestReallyAddHost(t *testing.T) {
	sut := inv.NewInv()
	group_name := "a group"
	host_name := "a host"
	sut.AddGroup(group_name)
	err := sut.AddHost(host_name, group_name)
	if err != nil {
		t.Fatal("there should be no error this time")
	}
	g, err := sut.GetGroup(group_name)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := g.Hosts[host_name]
	if !ok {
		t.Fatal("no host in group")
	}
}