package main

import "testing"

func TestHello(t *testing.T) {
	emoji := ":world_map:"
	hello := Hello(emoji)
	expected := "Hello, " + emoji

	if hello != expected {
		t.Errorf("hello %q expected %q", hello, expected)
	}
	t.Log("Success")
}
