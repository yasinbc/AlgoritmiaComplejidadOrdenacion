package ordenacion

func Insercion(array []int) []int {
	contador := 1

	for i, v := range array {
		if i < 1 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			contador++
			if array[j] <= v {
				break
			}
			array[i] = array[j]
			array[j] = v
			i--
		}
	}
	return array
}
