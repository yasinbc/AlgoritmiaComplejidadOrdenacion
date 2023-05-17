package main

import (
	"fmt"
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
	"time"
)

func main() {
	array := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	start := time.Now()
	fmt.Println(ordenacion.Burbuja(array))
	elapsed := time.Since(start)
	fmt.Println("==elapsed time==", elapsed)

}
