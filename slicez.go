// Package slicez provides utility functions for slices.
package slicez

// Remove unsafely removes the element at index.
func Remove[S ~[]E, E any](xs S, i int) S {
	if len(xs) == 0 {
		return nil
	}

	return append(xs[:i], xs[i+1:]...)
}

// Count counts the values of the slice that satisfy the condition.
func Count[S ~[]E, E any](xs S, fn func(val E) bool) int {
	var count int

	if len(xs) == 0 {
		return count
	}

	for i := range xs {
		if fn(xs[i]) {
			count++
		}
	}

	return count
}

// Unique returns a slice with all unique values from the slice.
func Unique[S ~[]E, E comparable](xs S) S {
	if len(xs) == 0 {
		return nil
	}

	var (
		unq = make(map[E]struct{})
		out = make(S, 0)
	)

	for i := range xs {
		if _, ok := unq[xs[i]]; !ok {
			unq[xs[i]] = struct{}{}

			out = append(out, xs[i])
		}
	}

	return out
}

// Difference returns a slice with all the values from the slice that are not in the others.
func Difference[S ~[]E, E comparable](xs S, slices ...S) S {
	if len(xs) == 0 {
		return nil
	}

	if len(slices) == 0 {
		return xs
	}

	var (
		exs = make(map[E]struct{})
		out = make(S, 0)
	)

	for i := range slices {
		for j := range slices[i] {
			exs[slices[i][j]] = struct{}{}
		}
	}

	for i := range xs {
		if _, ok := exs[xs[i]]; !ok {
			out = append(out, xs[i])
		}
	}

	return out
}

// Filter returns a slice with filtered values.
func Filter[S ~[]E, E any](xs S, fn func(val E) bool) S {
	if len(xs) == 0 {
		return nil
	}

	out := make(S, 0)

	for i := range xs {
		if fn(xs[i]) {
			out = append(out, xs[i])
		}
	}

	return out
}

// Map returns a slice with mapped values.
func Map[S ~[]E, E any](xs S, fn func(val E) E) S {
	if len(xs) == 0 {
		return nil
	}

	out := make(S, len(xs))

	for i := range xs {
		out[i] = fn(xs[i])
	}

	return out
}

// Remap returns a slice with remapped values.
func Remap[S ~[]E, C ~[]T, E, T any](xs S, fn func(val E) T) C {
	if len(xs) == 0 {
		return nil
	}

	out := make(C, len(xs))

	for i := range xs {
		out[i] = fn(xs[i])
	}

	return out
}

// Reduce reduces the values from the slice into the accumulator.
func Reduce[S ~[]E, E any](xs S, acc E, fn func(acc E, val E) E) E {
	if len(xs) == 0 {
		return acc
	}

	for i := range xs {
		acc = fn(acc, xs[i])
	}

	return acc
}

// ReduceDefault reduces the values from the slice into the accumulator.
//
// Accumulator value is the default type value.
func ReduceDefault[S ~[]E, E any](xs S, fn func(acc E, val E) E) E {
	var acc E

	return Reduce(xs, acc, fn)
}

// Flatten flattens the slice one time.
func Flatten[S ~[]E, E any](xs []S) S {
	if len(xs) == 0 {
		return nil
	}

	out := make(S, 0)

	for i := range xs {
		out = append(out, xs[i]...)
	}

	return out
}
