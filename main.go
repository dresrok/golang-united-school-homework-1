package main

import (
	"github.com/kyokomi/emoji/v2"
)

func Hello() (x string) {
	hello := "Hello :world_map:"
	return hello
}

func main() {
	emoji.Println(Hello())
}
