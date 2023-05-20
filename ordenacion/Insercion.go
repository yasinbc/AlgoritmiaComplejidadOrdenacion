package ordenacion

func Insercion(array []int) []int {
	// Contador de pasos
	stepCounter := 1

	// Evalua cada valor del slice
	for i, v := range array {
		//El primero no tiene valor que comparar, se omite
		if i < 1 {
			continue
		}

		//Recorre cada index hasta encontrar el valor neto al actual
		for j := i - 1; j >= 0; j-- {
			//fmt.Printf("%v > %v = %v\n", array[j], v, array[j] > v)
			stepCounter++

			if array[j] <= v {
				//fmt.Printf("%v\n", array)
				break
			}

			array[i] = array[j]
			array[j] = v
			i--

			//fmt.Printf("%v\n", array)
		}
	}
	return array
}
