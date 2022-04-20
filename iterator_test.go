package iterator_test

import (
	"strconv"
	"testing"

	"github.com/bartvsdev/iterator"
	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	iter := iterator.Slice(slice)
	for _, expected := range slice {
		actual, err := iter.Next()
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	}
	_, err := iter.Next()
	require.Error(t, err)
}

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3}
	iter := iterator.Slice(slice)
	mapIter := iterator.Map(iter, func(i int) int { return i + 1 })
	for _, original := range slice {
		mapped, err := mapIter.Next()
		require.NoError(t, err)
		require.Equal(t, original+1, mapped)
	}
	_, err := mapIter.Next()
	require.ErrorIs(t, err, iterator.ErrNoNext)
}

func TestMapErr(t *testing.T) {
	iter := iterator.Slice([]string{"1", "two", "3"})
	mapped := iterator.MapErr(iter, strconv.Atoi)
	// 1
	value, err := mapped.Next()
	require.NoError(t, err)
	require.Equal(t, 1, value)
	// two
	_, err = mapped.Next()
	require.ErrorIs(t, err, strconv.ErrSyntax)
	// 3
	value, err = mapped.Next()
	require.NoError(t, err)
	require.Equal(t, 3, value)
	// empty
	_, err = mapped.Next()
	require.ErrorIs(t, err, iterator.ErrNoNext)
}

func TestFilter(t *testing.T) {
	iter := iterator.Slice([]int{1, 2, 3, 4})
	filtered := iterator.Filter(iter, func(i int) bool { return i%2 == 0 })
	// 2
	value, err := filtered.Next()
	require.NoError(t, err)
	require.Equal(t, 2, value)
	// 4
	value, err = filtered.Next()
	require.NoError(t, err)
	require.Equal(t, 4, value)
	// empty
	_, err = filtered.Next()
	require.ErrorIs(t, err, iterator.ErrNoNext)
}

func TestForeach(t *testing.T) {
	expected := []int{1, 2, 3}
	iter := iterator.Slice(expected)
	var actual []int
	iterator.ForEach(iter, func(i int) { actual = append(actual, i) })
	require.Equal(t, expected, actual)
}

func TestFold(t *testing.T) {
	iter := iterator.Slice([]int{1, 2, 3})
	sum, err := iterator.Fold(iter, 36, func(a, i int) int { return a + i })
	require.NoError(t, err)
	require.Equal(t, 42, sum)
}

func TestToSlice(t *testing.T) {
	original := []int{1, 2, 3}
	iter := iterator.Slice(original)
	copy, err := iterator.ToSlice(iter)
	require.NoError(t, err)
	require.Len(t, copy, len(original))
	require.Equal(t, copy[0], original[0])
	require.Equal(t, copy[1], original[1])
	require.Equal(t, copy[2], original[2])
}
