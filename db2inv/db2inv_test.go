package db2inv_test

import(
	"testing"
	"yauai/db2inv"
	"yauai/db"
)

func TestAHost(t *testing.T) {
	in := []byte(`[{"foo":"bar"}]`)
	tdb, err := db.NewDb(in)
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