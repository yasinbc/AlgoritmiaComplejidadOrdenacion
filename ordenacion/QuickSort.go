package ordenacion

import (
	"math/rand"
)

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	izquierda, derecha := 0, len(array)-1

	pivote := rand.Int() % len(array)

	array[pivote], array[derecha] = array[derecha], array[pivote]

	for i, _ := range array {
		if array[i] < array[derecha] {
			array[izquierda], array[i] = array[i], array[izquierda]
			izquierda++
		}
	}

	array[izquierda], array[derecha] = array[derecha], array[izquierda]

	QuickSort(array[:izquierda])
	QuickSort(array[izquierda+1:])

	return array
}
