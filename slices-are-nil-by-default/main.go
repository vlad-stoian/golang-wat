package main

import (
	"fmt"
)

type Foo struct {
	AnInnocentSlice []string
}

func main() {
	fmt.Println(Foo{}.AnInnocentSlice == nil)
}
