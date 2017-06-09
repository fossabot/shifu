package main_test

import (
	"github.com/ayonious/shifu"
	"testing"
)

func TestAssertStr(t *testing.T) {
	main.AssertEquealString("stra","stra")
}

func TestAssertInt(t *testing.T) {
	main.AssertEquealInt(10,10)
}
