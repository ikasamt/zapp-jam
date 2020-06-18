package main

import (
	"testing"
)

func TestParseDir(t *testing.T) {
	pn, task, _ := ParseDir("./example/app")

	if pn != `app` {
		t.Fatalf("failed test %v", pn)
	}

	if task == nil {
		t.Fatalf("failed test %v", task)
	}

}
