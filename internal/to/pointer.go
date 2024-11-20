package to

func StringPtr[T ~string](s T) *T {
	return &s
}

func Uint8Ptr[T ~uint8](s T) *T { return &s }
