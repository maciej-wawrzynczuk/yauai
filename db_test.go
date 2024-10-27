package main

import (
	"testing"
)

func TestHosts(t *testing.T) {
	sut := db{
		{"foo": "bar"},
	}

	result := sut.hosts("foo")

	if result[0] != "bar" {
		t.Fatal("bad value")
	}
}