package slices

// Concat concatenates given slices into a single slice.
func Concat[S ~[]T, T any](slices ...S) S {
	size := 0
	for _, items := range slices {
		size += len(items)
	}
	result := make(S, 0, size)
	for _, items := range slices {
		result = append(result, items...)
	}
	return result
}

// Zip returns chan of arrays of elements
// from given arrays on the same position.
func Zip[S ~[]T, T any](items ...S) chan S {
	if len(items) == 0 {
		result := make(chan S)
		close(result)
		return result
	}

	size := len(items[0])
	for _, arr := range items[1:] {
		if len(arr) < size {
			size = len(arr)
		}
	}

	result := make(chan S, 1)
	go func() {
		for i := 0; i < size; i++ {
			chunk := make([]T, 0, len(items))
			for _, arr := range items {
				chunk = append(chunk, arr[i])
			}
			result <- chunk
		}
		close(result)
	}()
	return result
}
