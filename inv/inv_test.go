package inv_test

// TODO: test for _meta

import (
	"testing"
	"yauai/inv"
)


func TestAddGroup(t *testing.T) {
	group_name := "foo"
	sut := inv.NewInv()
	sut.AddGroup(group_name)

	if len(sut.Groups()) != 1{
		t.Fatal("Ther's should be one group")
	}
	
	_, err := sut.GetGroup(group_name)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddHostNoSuchGroup(t *testing.T) {
	sut := inv.NewInv()
	err := sut.AddHost("h", "g")
	if err == nil {
		t.Fatal("it should be an error")
	}
}

func TestReallyAddHost(t *testing.T) {
	sut := inv.NewInv()
	group_name := "a group"
	host_name := "a host"
	sut.AddGroup(group_name)
	err := sut.AddHost(host_name, group_name)
	if err != nil {
		t.Fatal(err)
	}
	g, err := sut.GetGroup(group_name)
	if err != nil {
		t.Fatal(err)
	}
	hosts := g.HostNames()
	good := false
	for _, h := range hosts {
		if h == host_name {
			good = true
			break
		}
	}
	if !good {
		t.Fatal("host not found")
	}
}

func TestAddGroupTwice(t *testing.T) {
	sut := inv.NewInv()
	g_name := "foo_group"
	h_name := "bar_host"
	sut.AddGroup(g_name)
	sut.AddHost(h_name, g_name)
	sut.AddGroup(g_name)

	g, err := sut.GetGroup(g_name)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, h := range g.HostNames() {
		if h == h_name {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("Host not found")
	}
}