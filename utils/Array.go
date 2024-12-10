package utils

func MapArray[Type any, Res any](arr []Type, transform func(Type) Res) []Res {
	result := make([]Res, len(arr))

	for el := range arr {
		result[el] = transform(arr[el])
	}

	return result
}
