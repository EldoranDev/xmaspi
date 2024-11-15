package to

func StringPtr[T ~string](s T) *T {
	return &s
}
