package main

import "testing"

func TestHello(t *testing.T) {
	hello := Hello()
	expected := "Hello :world_map:"

	if hello != expected {
		t.Errorf("hello %q expected %q", hello, expected)
	}
}
