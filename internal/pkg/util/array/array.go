package array

func GetOrDefault[T any](arr []T, index int, defaultValue T) T {
	if index >= 0 && index < len(arr) {
		return arr[index]
	}
	return defaultValue
}
