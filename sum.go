package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	result, err := processArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func processArgs(args []string) (string, error) {
	a, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(args[2])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d + %d = %d", a, b, sum(a, b)), nil
}

func sum(a, b int) int {
	return a + b
}
