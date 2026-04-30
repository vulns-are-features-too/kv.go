package database

func getMapKeys[T any](m map[string]T) []string {
	total := len(m)
	keys := make([]string, total)
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
