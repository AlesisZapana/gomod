package main

import "fmt"

// Funcion variádica
func suma(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func imprimirNombres(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(suma(1, 2, 3, 4))
	imprimirNombres("Aco", "Nabu", "bein")
	fmt.Println(getValues(2))
}
