package homework

func sortMapValues(input map[int]string) (result []string) {
	mn := -9999999
	for range input {
		mn2 := 99999999
		elemToAppend := ""
		for key, element := range input {
			if mn < key && mn2 > key {
				mn2 = key
				elemToAppend = element
			}
		}
		mn = mn2
		result = append(result, elemToAppend)
	}

	return result
}
