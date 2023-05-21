package ordenacion

func sort(A []int) []int {
	if len(A) <= 1 {
		return A
	}

	left, right := split(A)
	left = sort(left)
	right = sort(right)
	return merge(left, right)
}

// split array into two
func split(A []int) ([]int, []int) {
	return A[0 : len(A)/2], A[len(A)/2:]
}

// assumes that A and B are sorted
func merge(A, B []int) []int {
	arr := make([]int, len(A)+len(B))

	// index j for A, k for B
	j, k := 0, 0

	for i := 0; i < len(arr); i++ {
		// fix for index out of range without using sentinel
		if j >= len(A) {
			arr[i] = B[k]
			k++
			continue
		} else if k >= len(B) {
			arr[i] = A[j]
			j++
			continue
		}
		// default loop condition
		if A[j] > B[k] {
			arr[i] = B[k]
			k++
		} else {
			arr[i] = A[j]
			j++
		}
	}

	return arr
}

//func MergeSort(items []int) []int {
//	if len(items) < 2 {
//		return items
//	}
//	first := MergeSort(items[:len(items)/2])
//	second := MergeSort(items[len(items)/2:])
//	return Merge(first, second)
//}
//
//func Merge(a []int, b []int) []int {
//	final := []int{}
//	i := 0
//	j := 0
//	for i < len(a) && j < len(b) {
//		if a[i] < b[j] {
//			final = append(final, a[i])
//			i++
//		} else {
//			final = append(final, b[j])
//			j++
//		}
//	}
//	for ; i < len(a); i++ {
//		final = append(final, a[i])
//	}
//	for ; j < len(b); j++ {
//		final = append(final, b[j])
//	}
//	return final
//}

//func MergeSort(array []int) []int {
//	if len(array) < 2 {
//		return array
//	}
//
//	mitad := (len(array)) / 2
//	return Merge(MergeSort(array[:mitad]), MergeSort(array[mitad:]))
//
//}
//
//func Merge(izquierda, derecha []int) []int {
//
//	size, i, j := len(izquierda)+len(derecha), 0, 0
//	array := make([]int, size, size)
//
//	for k := 0; k < size; k++ {
//		if i > len(izquierda)-1 && j <= len(derecha)-1 {
//			array[k] = derecha[j]
//			j++
//		} else if j > len(derecha)-1 && i <= len(izquierda)-1 {
//			array[k] = izquierda[i]
//			i++
//		} else if izquierda[i] < derecha[j] {
//			array[k] = izquierda[i]
//			i++
//		} else {
//			array[k] = derecha[j]
//			j++
//		}
//	}
//	return array
//}
