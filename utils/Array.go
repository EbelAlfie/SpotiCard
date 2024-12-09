package utils

func MapArray[Type any, Res any](arr []Type, transform func(Type) Res) []Res {
	result := make([]Res, len(arr)-1)

	for el := range arr {
		result = append(result, transform(arr[el]))
	}

	return result
}
