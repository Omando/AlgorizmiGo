package recursion

func GrayCode(n int) []string {
	var codes []string

	gray(n-1, []byte{'0'}, &codes)
	yarg(n-1, []byte{'1'}, &codes)

	return codes
}

func gray(n int, code []byte, codes *[]string) {
	// Exit conditions
	if n == 0 {
		*codes = append(*codes, string(code))
		return
	}

	code0 := append(code, '0')
	gray(n-1, code0, codes)
	code1 := append(code, '1')
	yarg(n-1, code1, codes)
}

func yarg(n int, code []byte, codes *[]string) {
	// Exit conditions
	if n == 0 {
		*codes = append(*codes, string(code))
		return
	}

	code0 := append(code, '1')
	gray(n-1, code0, codes)
	code1 := append(code, '0')
	yarg(n-1, code1, codes)
}
