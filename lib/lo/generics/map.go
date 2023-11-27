package generics

// Values creates an array of the map values.
// Play: https://go.dev/play/p/nnRTQkzQfF6
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))

	for _, v := range in {
		result = append(result, v)
	}

	return result
}
