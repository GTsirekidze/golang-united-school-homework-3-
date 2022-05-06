package homework

func average(input [15]float32) (result float32) {

	for i := range input {
		result += input[i]
	}

	return result / 15
}
