package main

import (
	"testing"
)

func TestParseDir(t *testing.T) {
	pn, task := ParseDir("./example/app")

	if pn != `app` {
		t.Fatal("failed test")
	}

	if task == nil {
		t.Fatal("failed test")
	}

}
