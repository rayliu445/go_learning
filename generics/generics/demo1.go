package generics

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}

	}
	return -1
}
