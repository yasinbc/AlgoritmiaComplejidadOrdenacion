package main

import (
	"github.com/yasinbc/AlgoritmiaComplejidad/ordenacion"
)

func main() {
	//array1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 12}
	//array2 := []int{-356, 328, 705, -199, -373, 108, -377, -362, 128, 98, 1, -9, -500, -607, 387, 12, 210, -600, -351, 432}
	array3 := []int{2, 56843, -76716, -42042, 30846, 50290, -27458, 47259, 91477, 6775, 59768, 15476, -21788, -33386, 42947, -27244, -45848, -7689, 14298, -41118, -99082, -3641, -27944, -59233, 25235, -44990, 56287, -13288, 8899, 2579, -28954, 27772, 41360, 34462, 8174, 53027, -22394, -91773, -42887, 14202, 30264, 42811, -43774, -3273, 11149, 27266, -32695, 31350, -5934, 30010, -32840}

	ordenacion.TiempoEjecucion("Burbuja", ordenacion.Burbuja(array3))
	ordenacion.TiempoEjecucion("Selección", ordenacion.Seleccion(array3))
	ordenacion.TiempoEjecucion("Merge Sort", ordenacion.MergeSort(array3))
}
