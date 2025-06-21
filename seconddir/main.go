package seconddir

func ReverseString(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

