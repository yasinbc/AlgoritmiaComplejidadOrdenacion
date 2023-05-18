package ordenacion

import (
	"fmt"
	"time"
)

func TiempoEjecucion(nombre string, ordenacion []int) {
	fmt.Println("-------------------------------")
	startTime := time.Now()
	fmt.Println(nombre + ": ")
	fmt.Println(ordenacion)
	endTime := time.Now()
	fmt.Println("Tiempo al empezar: ", startTime)
	fmt.Println("Tiempo al terminar: ", endTime)
	fmt.Println("-------------------------------")
}