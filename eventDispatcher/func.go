package eventDispatcher


func InSliceString(t[]string, s string) bool {
	for _, v := range t {
		if v == s {
			return true
		}
	}
	return false
}

func AddString(t[]string, s string) []string {
	if	InSliceString(t, s) {
		return t
	}
	return append(t, s)
}

func RemoveString(t[]string, s string) []string {
	for i := 0; i < len(t); i++ {
		if t[i] == s {
			t = append(t[:i], t[i+1:]...)
			i--
		}
	}

	return t
}
