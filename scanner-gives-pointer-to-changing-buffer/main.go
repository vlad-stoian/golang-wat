package main

import (
	"bufio"
	"fmt"
  "os"
)

type Line []byte

func main() {
	file, err := os.Open("testfile")
	if err != nil {
    panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 10)
	scanner.Buffer(buf, 10)

  var lines []Line
  for scanner.Scan() {
    lines = append(lines, scanner.Bytes())
  }

  //MADNESS!!!
  for _, line := range lines {
    fmt.Println(string(line))
  }
}
