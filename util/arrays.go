package util

func Contains[T int | int8 | int16 | int32 | int64 | string](item T, Arr []T) bool {
	for _, v := range Arr {
		if v == item {
			return true
		}
	}
	return false
}
