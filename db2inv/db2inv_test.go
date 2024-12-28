package db2inv_test

import (
	"testing"
	"yauai/db"
	"yauai/db2inv"
)

func TestAHost(t *testing.T) {
	in := []byte(`[{"foo":"bar"}]`)
	tdb, err := db.FromJson(in)
	if err != nil {
		t.Fatal("it shouldn't happen!!!!")
	}

	sut, err := db2inv.Db2inv(*tdb, "foo")
	if err != nil {
		t.Fatal(err)
	}

	g, err := sut.GetGroup("ungroupped")
	if err != nil {
		t.Fatal(err)
	}
	names := g.HostNames()
	if names[0] != "bar" {
		t.Fatal("bad hostname")
	}
}

func Test1group(t *testing.T) {
	data := []byte(`[{"hostname": "host1", "group": "group1"}]`)
	_, err := db.FromJson(data)
	if err != nil {
		t.Fatal(err)
	}
}
