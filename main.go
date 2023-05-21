package main

import (
	"fmt"
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
	"time"
)

func main() {
	//Incluye visualizacion de las matrices
	var size int = 100
	array := ordenacion.RanNumArr(size)

	fmt.Println("\n\n\n\nArray de ", size, " elementos a ordenar.\n")
	tiempoPrograma := time.Now()

	fmt.Println(array, "\n")
	fmt.Println("	- Algoritmo de Ordenacion Burbuja: ")
	startTime1 := time.Now()
	ordenacion.Burbuja(array)
	fmt.Println("		Tiempo de ejecucion: ")
	fmt.Println("			Segundos:      ", time.Since(startTime1).Seconds())
	fmt.Println("			Milisegundos:  ", time.Since(startTime1).Milliseconds())
	fmt.Println("			Microsegundos: ", time.Since(startTime1).Microseconds())
	fmt.Println("			Nanosegundos:  ", time.Since(startTime1).Nanoseconds(), "\n")

	array = nil
	array = ordenacion.RanNumArr(size)

	fmt.Println("	- Algoritmo de Ordenacion por Seleccion: ")
	startTime2 := time.Now()
	ordenacion.Seleccion(array)
	fmt.Println("		Tiempo de ejecucion: ")
	fmt.Println("			Segundos:      ", time.Since(startTime2).Seconds())
	fmt.Println("			Milisegundos:  ", time.Since(startTime2).Milliseconds())
	fmt.Println("			Microsegundos: ", time.Since(startTime2).Microseconds())
	fmt.Println("			Nanosegundos:  ", time.Since(startTime2).Nanoseconds(), "\n")

	array = nil
	array = ordenacion.RanNumArr(size)

	fmt.Println("	- Algoritmo de Ordenacion por Insercion: ")
	startTime3 := time.Now()
	ordenacion.Insercion(array)
	fmt.Println("		Tiempo de ejecucion: ")
	fmt.Println("			Segundos:      ", time.Since(startTime3).Seconds())
	fmt.Println("			Milisegundos:  ", time.Since(startTime3).Milliseconds())
	fmt.Println("			Microsegundos: ", time.Since(startTime3).Microseconds())
	fmt.Println("			Nanosegundos:  ", time.Since(startTime3).Nanoseconds(), "\n")

	array = nil
	array = ordenacion.RanNumArr(size)

	fmt.Println("	- Algoritmo de Ordenacion Merge Sort: ")
	startTime4 := time.Now()
	ordenacion.MergeSort(array, 0, len(array)-1)
	fmt.Println("		Tiempo de ejecucion: ")
	fmt.Println("			Segundos:      ", time.Since(startTime4).Seconds())
	fmt.Println("			Milisegundos:  ", time.Since(startTime4).Milliseconds())
	fmt.Println("			Microsegundos: ", time.Since(startTime4).Microseconds())
	fmt.Println("			Nanosegundos:  ", time.Since(startTime4).Nanoseconds(), "\n")

	array = nil
	array = ordenacion.RanNumArr(size)

	fmt.Println("	- Algoritmo de Ordenacion Quick Sort: ")
	startTime5 := time.Now()
	ordenacion.QuickSort(array)
	fmt.Println("		Tiempo de ejecucion: ")
	fmt.Println("			Segundos:      ", time.Since(startTime5).Seconds())
	fmt.Println("			Milisegundos:  ", time.Since(startTime5).Milliseconds())
	fmt.Println("			Microsegundos: ", time.Since(startTime5).Microseconds())
	fmt.Println("			Nanosegundos:  ", time.Since(startTime5).Nanoseconds(), "\n")

	fmt.Println("Tiempo de ejecucion de todo el programa: ")
	fmt.Println("	Segundos:      ", time.Since(tiempoPrograma).Seconds())
	fmt.Println("	Milisegundos:  ", time.Since(tiempoPrograma).Milliseconds())
	fmt.Println("	Microsegundos: ", time.Since(tiempoPrograma).Microseconds())
	fmt.Println("	Nanosegundos:  ", time.Since(tiempoPrograma).Nanoseconds(), "\n")

}
