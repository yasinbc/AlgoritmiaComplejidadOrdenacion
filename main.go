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

	//fmt.Println("Algoritmo de Ordenacion por Seleccion: ")
	//startTime2 := time.Now()
	//ordenacion.Seleccion(array)
	//fmt.Println("Tiempo de ejecucion: ", time.Since(startTime2))
	//fmt.Println("\n")
	//
	//fmt.Println("Algoritmo de Ordenacion por Insercion: ")
	//startTime3 := time.Now()
	//ordenacion.Insercion(array)
	//fmt.Println("Tiempo de ejecucion: ", time.Since(startTime3))
	//fmt.Println("\n")

	tiempoPrograma := time.Now()

	fmt.Println("\n\n\nAlgoritmo de Ordenacion Burbuja: ")
	startTime1 := time.Now()
	ordenacion.Burbuja(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime1))

	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion por Seleccion: ")
	startTime2 := time.Now()
	ordenacion.Seleccion(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime2))
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion por Insercion: ")
	startTime3 := time.Now()
	ordenacion.Insercion(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime3))
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion Merge Sort: ")
	startTime4 := time.Now()
	ordenacion.MergeSort(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime4))
	fmt.Println("\n")

	fmt.Println("Algoritmo de Ordenacion Quick Sort: ")
	startTime5 := time.Now()
	ordenacion.QuickSort(array)
	fmt.Println("Tiempo de ejecucion: ", time.Since(startTime5))
	fmt.Println("\n")

	fmt.Println("\n\n\n\nTiempo de ejecucion de todo el programa: ", time.Since(tiempoPrograma))
	//ordenacion.TiempoEjecucion("Burbuja", ordenacion.Burbuja(array3))
	//ordenacion.TiempoEjecucion("Selección", ordenacion.Seleccion(array3))
	//ordenacion.TiempoEjecucion("Merge Sort", ordenacion.MergeSort(array3))
	//ordenacion.TiempoEjecucion("Insercion", ordenacion.Insercion(array3))
	//ordenacion.TiempoEjecucion("QuickSort", ordenacion.QuickSort(array3))

}
