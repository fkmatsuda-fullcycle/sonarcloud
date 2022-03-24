package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d + %d = %d\n", a, b, sum(a, b))
}

func sum(a, b int) int {
	return a + b
}
