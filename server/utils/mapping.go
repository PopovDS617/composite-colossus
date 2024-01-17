package utils

func SliceToInterface[T any](dataSlice []T) []interface{} {

	result := make([]interface{}, len(dataSlice))

	for i, v := range dataSlice {
		result[i] = v
	}
	return result
}
