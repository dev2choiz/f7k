package utils

type slice struct {}

func (r *slice) DeduplicateStrings(str []string) []string {
	//@todo : try to use r.Deduplicate() instead of ...
	m := make(map[string]bool)
	n := make([]string, 0)
	for _, s := range str {
		if _, v := m[s]; !v {
			n = append(n, s)
			m[s] = true
		}
	}

	return n
}

func (r *slice) Deduplicate(str []interface{}) []interface{} {
	m := make(map[interface{}]bool)
	n := make([]interface{}, 0)
	for _, s := range str {
		if _, v := m[s]; !v {
			n = append(n, s)
			m[s] = true
		}
	}

	return n
}
