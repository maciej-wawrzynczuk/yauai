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
//	any, _, err := testutils.Extract(*sut)
//	if err != nil {
//		t.Fatal(err)
//	}
//	_, ok := (*any)[group_name]
//	if !ok {
//		t.Fail()
//	}
}

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
		t.Fatal(err)
	}
//	any, _, err := testutils.Extract(*sut)
//	if err != nil {
//		t.Fatal(err)
//	}
//	g, ok := (*any)[group_name]
//	if !ok {
//		t.Fatal("no group")
//	}
//
//	h, ok := g["hosts"]
//	if !ok {
//		t.Fatal("no hosts entry")
//	}
//	_, ok = h[host_name]
//	if !ok {
//		t.Fatal("no host")
//	}
}

func TestAddGroupTwice(t *testing.T) {
	sut := inv.NewInv()
	g_name := "foo_group"
	h_name := "bar_host"
	sut.AddGroup(g_name)
	sut.AddHost(h_name, g_name)
	sut.AddGroup(g_name)

//	any, _, err := testutils.Extract(*sut)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	g, ok := (*any)[g_name]
//	if !ok {
//		t.Fatal("Group not found")
//	}
//	hosts, ok := g["hosts"]
//	if !ok {
//		t.Fatal("no hosts field. Really nasty")
//	}
//	_, ok = hosts[h_name]
//	if !ok {
//		t.Fatal("host not found")
//	}
}

func TestEmptyMeta(t *testing.T) {
//	sut := inv.NewInv()
//	_, _, err := testutils.Extract(*sut)
//	if err != nil {
//		t.Fatal(err)
//	}
	
}