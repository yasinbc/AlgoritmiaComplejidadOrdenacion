package ordenacion

func Sort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	izquierda, derecha := split(array)
	izquierda = Sort(izquierda)
	derecha = Sort(derecha)
	return merge(izquierda, derecha)
}

// Divide el array en 2
func split(array []int) ([]int, []int) {
	return array[0 : len(array)/2], array[len(array)/2:]
}

// Asume que arrayA y arrayB estan ordenadas
func merge(arrayA, arrayB []int) []int {
	arr := make([]int, len(arrayA)+len(arrayB))

	// index j para arrayA, k para arrayB
	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		if j >= len(arrayA) {
			arr[i] = arrayB[k]
			k++
			continue
		} else if k >= len(arrayB) {
			arr[i] = arrayA[j]
			j++
			continue
		}

		if arrayA[j] > arrayB[k] {
			arr[i] = arrayB[k]
			k++
		} else {
			arr[i] = arrayA[j]
			j++
		}
	}

	return arr
}
