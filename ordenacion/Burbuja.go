package ordenacion

func Burbuja(arrayDesordenada []int) []int {
	for i := 0; i < len(arrayDesordenada)-1; i++ {
		for j := 0; j < len(arrayDesordenada)-i-1; j++ {
			if arrayDesordenada[j] > arrayDesordenada[j+1] {
				arrayDesordenada[j], arrayDesordenada[j+1] = arrayDesordenada[j+1], arrayDesordenada[j]
			}
		}
	}
	return arrayDesordenada
}
