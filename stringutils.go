package commonlibs

func Substring(in string, start, length int) string {
	size := len(in)
	endIndex := start + length

	if length == 0 || start < 0 || start > size || endIndex > size {
		return ""
	}

	return in[start:endIndex]
}

func SubstringInPlace(in *string, start, length int) {
	size := len(*in)
	endIndex := start + length

	if length == 0 || start < 0 || start > size || endIndex > size {
		return
	}

	*in = (*in)[start:endIndex]
}
