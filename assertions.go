package main


import (
	"fmt"
	"os"
)


func AssertEquealString(actual string, expected string) {
	if actual != expected {
		fmt.Printf("❌ didnt match expected result, expected: %v ,found: %v\n" , expected, actual)
		os.Exit(1)
	} else {
		fmt.Printf("✅ matched result: %v\n", expected)
	}
}

func AssertEquealInt(actual int, expected int) {
	if actual != expected {
		fmt.Printf("❌ didnt match expected result, expected: %v ,found: %v\n" , expected, actual)
		os.Exit(1)
	} else {
		fmt.Printf("✅ matched result: %v\n", expected)
	}
}
