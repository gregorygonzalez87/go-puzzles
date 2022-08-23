package main

import (
	"fmt"
	"strings"
)

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		spaces := strings.Repeat(" ", int(n)-1-i)
		stairs := strings.Repeat("#", i+1)
		fmt.Printf("%s%s\n", spaces, stairs)
	}
}

func main() {
	staircase(6)
}
