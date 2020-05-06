package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Panicf("unexpected args %s", os.Args)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panicf("unexpected number %s", os.Args[1])
	}

	fmt.Printf("n [%d], fib [%d]\n", n, fib(uint64(n)))
}
