package solution

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	hello := GetMessage()
	expected := emoji.Sprint("Hello :world_map:!")

	if hello != expected {
		t.Errorf("hello %q expected %q", hello, expected)
	}
}
