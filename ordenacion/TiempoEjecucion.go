package ordenacion

import (
	"fmt"
	"time"
)

func TiempoEjecucion(nombre string) func() {
	fmt.Println("")
	fmt.Println("-------------------------------")

	startTime := time.Now()
	return func() {

		fmt.Println("s% Tardó %v\n", nombre, time.Since(startTime))
	}
	//fmt.Println("-------------------------------")
	//fmt.Println("")
}
