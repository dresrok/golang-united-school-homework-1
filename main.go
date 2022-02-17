package main

import (
	"github.com/kyokomi/emoji/v2"
)

func Hello(emoji string) string {
	return "Hello, " + emoji
}

func main() {
	emoji.Println(Hello(":world_map:"))
}
