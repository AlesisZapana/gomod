package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)
	z := 10
	w := func() int {
		return z * 2
	}()
	fmt.Println(w)

	c := make(chan int)
	go func() {
		fmt.Println("Iniciar Función")
		time.Sleep(5 * time.Second)
		fmt.Println("Fin")
		c <- 1
	}()
	<-c
}
