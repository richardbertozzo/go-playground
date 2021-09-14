package main

import (
	"fmt"
)

func main() {
	panicSystem()
}

func panicSystem() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("test")
}
