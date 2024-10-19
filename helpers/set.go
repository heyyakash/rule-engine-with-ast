package helpers

func GenerateSet(rules []string) []string {
	m := map[string]bool{}
	arr := []string{}
	for _, v := range rules {
		if _, ok := m[v]; !ok {
			arr = append(arr, v)
			m[v] = true
		}
	}
	return arr
}
