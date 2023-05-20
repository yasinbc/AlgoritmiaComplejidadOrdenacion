package main

import (
	"fmt"
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
	"time"
)

func main() {

	//defer ordenacion.TiempoEjecucion("Burbuja")()
	//time.Sleep(time.Second * 2)

	//longitud del array
	var size int = 100000
	array := ordenacion.RanNumArr(size)

	fmt.Println("Algoritmo de Ordenacion Burbuja: \n")
	startTime1 := time.Now()
	ordenacion.Burbuja(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime1))
	//endTime1 := time.Now()
	//fmt.Println("Tiempo al iniciar: ", startTime1)
	//fmt.Println("Tiempo al terminar: ", endTime1)
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion por Seleccion: \n")
	startTime2 := time.Now()
	ordenacion.Seleccion(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime2))
	//endTime2 := time.Now()
	//fmt.Println("Tiempo al iniciar: ", startTime2)
	//fmt.Println("Tiempo al terminar: ", endTime2)
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion por Insercion: \n")
	startTime3 := time.Now()
	ordenacion.Insercion(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime3))
	//endTime3 := time.Now()
	//fmt.Println("Tiempo al iniciar: ", startTime3)
	//fmt.Println("Tiempo al terminar: ", endTime3)
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion Merge Sort: \n")
	startTime4 := time.Now()
	ordenacion.MergeSort(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime4))
	//endTime4 := time.Now()
	//fmt.Println("Tiempo al iniciar: ", startTime4)
	//fmt.Println("Tiempo al terminar: ", endTime4)
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion Quick Sort: \n")
	startTime5 := time.Now()
	ordenacion.QuickSort(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime5))
	//endTime5 := time.Now()
	//fmt.Println("Tiempo al iniciar: ", startTime5)
	//fmt.Println("Tiempo al terminar: ", endTime5)
	fmt.Println("\n")

	//ordenacion.TiempoOrdenacion(ordenacion.Burbuja(array3))

	//ordenacion.TiempoEjecucion("Burbuja", ordenacion.Burbuja(array3))
	//ordenacion.TiempoEjecucion("Selección", ordenacion.Seleccion(array3))
	//ordenacion.TiempoEjecucion("Merge Sort", ordenacion.MergeSort(array3))
	//ordenacion.TiempoEjecucion("Insercion", ordenacion.Insercion(array3))
	//ordenacion.TiempoEjecucion("QuickSort", ordenacion.QuickSort(array3))

}
