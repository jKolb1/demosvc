package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("hello world... %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
