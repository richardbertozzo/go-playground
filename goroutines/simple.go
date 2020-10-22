package main

import (
	"fmt"
	"time"
)

type Product struct {
	Name     string
	Quantity int
}

func main() {
	p := Product{Name: "Coffe", Quantity: 10}
	go UpdateStock(p, true)

	fmt.Println("done")
	time.Sleep(time.Second) // espera 1 sec, para a goroutine anterior executar antes que o runtime termine
}

func UpdateStock(p Product, increment bool) {
	if increment {
		p.Quantity++
	} else {
		p.Quantity--
	}
}
