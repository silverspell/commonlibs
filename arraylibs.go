package commonlibs

func Map[T any](input []T, f func(T) T) []T {
	var result []T
	for _, v := range input {
		result = append(result, f(v))
	}
	return result
}

func MapInPlace[T any](input *[]T, f func(T) T) {
	for i, v := range *input {
		(*input)[i] = f(v)
	}
}

func Filter[T any](input []T, f func(T) bool) []T {
	var result []T
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterInPlace[T any](input *[]T, f func(T) bool) {
	n := 0
	//y := *input
	for _, v := range *input {
		if f(v) {
			(*input)[n] = v
			n += 1
		}
	}
	*input = (*input)[:n]
}

func Reduce[T, K any](input []T, initial K, f func(K, T) K) K {
	result := initial
	for _, v := range input {
		result = f(result, v)
	}
	return result
}

func Reverse[T any](input *[]T) {
	x := *input
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	*input = x
}

func Any[T any](input *[]T, f func(T) bool) bool {
	for _, v := range *input {
		if f(v) {
			return true
		}
	}
	return false
}

func Every[T any](input *[]T, f func(T) bool) bool {
	for _, v := range *input {
		if !f(v) {
			return false
		}
	}
	return true
}

func Find[T any](input *[]T, f func(T) (bool, *T)) (bool, *T) {
	for _, v := range *input {

		if ok, val := f(v); ok {
			return true, val
		}
	}
	return false, nil
}
