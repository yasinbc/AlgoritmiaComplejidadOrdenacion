package main

import (
	"fmt"
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
	"time"
)

func main() {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	startTime0 := time.Now()
	fmt.Println("Burbuja: ")
	fmt.Println(ordenacion.Burbuja(array))
	endTime0 := time.Now()
	fmt.Println("Tiempo al empezar: ", startTime0)
	fmt.Println("Tiempo al terminar: ", endTime0)

	startTime1 := time.Now()
	fmt.Println("Burbuja: ")
	fmt.Println(ordenacion.Seleccion(array))
	endTime1 := time.Now()
	fmt.Println("Tiempo al empezar: ", startTime1)
	fmt.Println("Tiempo al terminar: ", endTime1)
}
