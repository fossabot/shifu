package shifu

import (
	"github.com/ayonious/shifu"
	"testing"
)

func TestAssertStr(t *testing.T) {
	shifu.AssertEquealString("stra","stra")
}

func TestAssertInt(t *testing.T) {
	shifu.AssertEquealInt(10,10)
}
