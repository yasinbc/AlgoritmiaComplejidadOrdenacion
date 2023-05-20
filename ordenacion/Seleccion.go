package ordenacion

func Seleccion(arrayDesordenada []int) []int {

	contador := 1

	for i := 0; i < len(arrayDesordenada)-1; i++ {
		menor := i
		for j := i + 1; j < len(arrayDesordenada); j++ {
			contador++
			if arrayDesordenada[menor] > arrayDesordenada[j] {
				menor = j
			}
		}

		v := arrayDesordenada[i]
		arrayDesordenada[i] = arrayDesordenada[menor]
		arrayDesordenada[menor] = v
	}

	return arrayDesordenada
}
