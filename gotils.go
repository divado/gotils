package gotils

func reduce[T, M any](s []T, f func(M, T) M, initialValue M) M {
	accumulator := initialValue
	for _, v := range s {
		accumulator = f(accumulator, v)
	}
	return accumulator
}
