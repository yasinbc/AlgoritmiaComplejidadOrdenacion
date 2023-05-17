package main

import (
	"fmt"
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
	"time"
)

func main() {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	startTime := time.Now()
	fmt.Println(ordenacion.Burbuja(array))
	endTime := time.Now()
	fmt.Println("Tiempo al empezar: ", startTime)
	fmt.Println("Tiempo al terminar: ", endTime)

}
