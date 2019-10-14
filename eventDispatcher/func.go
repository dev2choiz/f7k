package eventDispatcher


func RemoveString(t[]string, s string) []string {
	for i := 0; i < len(t); i++ {
		if t[i] == s {
			t = append(t[:i], t[i+1:]...)
			i--
		}
	}

	return t
}
