package helpers

func IndexOf[S comparable](stack []S, needle S) (index int) {
	for index, value := range stack {
		if value == needle {
			return index
		}
	}
	return -1
}
