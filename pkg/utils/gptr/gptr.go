package gptr

func Of[T any](v T) *T {
	return &v
}

func Indirect[T any](p *T) T {
	if p != nil {
		return *p // Dereference p to get the value it points to.
	}
	var zero T  // Create a zero value of type T.
	return zero // Return the zero value.
}
