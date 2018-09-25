package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	ThisIsBool bool `json:"thisIsBool,omitempty"`
}

func main() {
	x, _ := json.Marshal(Foo{ThisIsBool: false})
	fmt.Println(string(x))
}
