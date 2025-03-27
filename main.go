package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() { ch <- "Hello from another thread!" }()
	fmt.Println(<-ch)
}
