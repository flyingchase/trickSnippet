package deferDemo

import (
	"fmt"
)

func Demo1() (index int) {
	defer fmt.Printf("defer fp(%v)\n", index)

	defer func() {
		fmt.Printf("defer func() {fp(%v)} ()\n", index)
		index++
	}()

	index++

	return
}
