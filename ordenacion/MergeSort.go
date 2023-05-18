package ordenacion

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mitad := len(array) / 2
	izquierda := MergeSort(array[:mitad])
	derecha := MergeSort(array[mitad:])
	return Merge(izquierda, derecha)

}

func Merge(izquierda, derecha []int) []int {
	resultado := make([]int, len(izquierda)+len(derecha))
	i, j := 0, 0

	for k := 0; k < len(resultado); k++ {
		if i >= len(izquierda) {
			resultado[k] = derecha[j]
			j++
		} else if j >= len(derecha) {
			resultado[k] = izquierda[i]
			i++
		} else if izquierda[i] < derecha[j] {
			resultado[k] = izquierda[i]
			i++
		} else {
			resultado[k] = derecha[j]
			j++
		}
	}
	return resultado
}
