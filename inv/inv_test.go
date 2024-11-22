package inv_test

// TODO: better testing. Output does not show a host despite aparently proper data
// don't test private stuff

import (
	"encoding/json"
	"testing"
	"yauai/inv"
)

func TestAddGroup(t *testing.T) {
	group_name := "foo"
	sut := inv.NewInv()
	sut.AddGroup(group_name)

	text, err := json.Marshal(sut)
	if err != nil {
		t.Fatal(err)
	}

	var any interface{}
	err = json.Unmarshal(text, &any)
	if err != nil {
		t.Fatal(err)
	}

	a_map, ok := any.(map[string]interface{})
	if !ok {
		t.Fail()
	}
	_, ok = a_map[group_name]
	if !ok {
		t.Fail()
	}
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
		t.Fatal("there should be no error this time")
	}
	g, err := sut.GetGroup(group_name)
	if err != nil {
		t.Fatal(err)
	}
	
	if (g.HostNames())[0] != host_name {
		t.Fatal("wrong hostname")
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
		t.Fatal("no group")
	}
	if len(g.HostNames()) != 1 {
		t.Fatal("there should be exactly one")
	}
}