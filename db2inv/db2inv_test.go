package db2inv_test

import (
	"slices"
	"testing"
	"yauai/db"
	"yauai/db2inv"
	// "yauai/inv"
)

func TestAHost(t *testing.T) {
	in := []byte(`[{"foo":"bar"}]`)
	tdb, err := db.FromJson(in)
	if err != nil {
		t.Fatal("it shouldn't happen!!!!")
	}

	sut := db2inv.New("foo")
	e := (*tdb)[0]
	err = sut.AddEntry(e)
	if err != nil {
		t.Fatal(err)
	}

	g, err := sut.Inv().GetGroup("ungroupped")
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(g.HostNames(), "bar") {
		t.Fatal("no host!")
	}
}

func Test1group(t *testing.T) {
	data := []byte(`[{"hostname": "host1", "group": "group1"}]`)
	db, err := db.FromJson(data)
	if err != nil {
		t.Fatal(err)
	}

	sut, err := db2inv.MkInv(db, "hostname", "group")
	if err != nil {
		t.Fatal(err)
	}

	g, err := sut.GetGroup("group1")
	if err != nil {
		t.Fatal(err)
	}

	if !slices.Contains(g.HostNames(), "host1") {
		t.Fatal("No host")
	}
}
