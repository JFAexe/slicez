package slicez_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JFAexe/slicez"
)

func Test_Slicez(t *testing.T) {
	t.Parallel()

	var (
		xs = []int{1, 2, 3, 4, 5, 6}
		ys = []int{1, 2, 2, 3, 4, 4, 5, 6}
		os = []int{1, 3, 5}
		es = []int{2, 4, 6}
		ss = []string{"1", "2", "3", "4", "5", "6"}
	)

	var (
		valTrue   = func(val int) bool { return true }
		valEven   = func(val int) bool { return val%2 == 0 }
		valPlus   = func(val int) int { return val + 1 }
		valString = func(val int) string { return strconv.Itoa(val) }
		accSum    = func(acc int, val int) int { return acc + val }
	)

	t.Run("Remove", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Remove[[]int](nil, 1))
		require.Equal(t, []int{1, 3, 4, 5, 6}, slicez.Remove(append([]int(nil), xs...), 1))
		require.Panics(t, func() { slicez.Remove(xs, 7) })
	})

	t.Run("Count", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, 0, slicez.Count[[]int](nil, valTrue))
		require.Equal(t, 3, slicez.Count(xs, valEven))
	})

	t.Run("Unique", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Unique[[]int](nil))
		require.Equal(t, xs, slicez.Unique(xs))
		require.Equal(t, xs, slicez.Unique(ys))
	})

	t.Run("Difference", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Difference[[]int](nil))
		require.Equal(t, xs, slicez.Difference(xs))
		require.Equal(t, es, slicez.Difference(xs, os))
	})

	t.Run("Filter", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Filter[[]int](nil, valTrue))
		require.Equal(t, es, slicez.Filter(xs, valEven))
	})

	t.Run("Map", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Map[[]int](nil, valPlus))
		require.Equal(t, es, slicez.Map(os, valPlus))
	})

	t.Run("Remap", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Remap[[]int, []string](nil, valString))
		require.Equal(t, ss, slicez.Remap[[]int, []string](xs, valString))
	})

	t.Run("Reduce", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, 42, slicez.Reduce[[]int](nil, 42, accSum))
		require.Equal(t, 42, slicez.Reduce(xs, 21, accSum))
	})

	t.Run("ReduceDefault", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, 0, slicez.ReduceDefault[[]int](nil, accSum))
		require.Equal(t, 21, slicez.ReduceDefault(xs, accSum))
	})

	t.Run("Flatten", func(t *testing.T) {
		t.Parallel()

		require.Nil(t, slicez.Flatten[[]int](nil))
		require.Equal(t, append(os, es...), slicez.Flatten([][]int{os, es}))
	})
}
