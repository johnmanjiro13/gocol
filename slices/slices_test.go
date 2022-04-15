package slices_test

import (
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"

	"johnmanjiro13/gocol/slices"
)

func TestMap(t *testing.T) {
	t.Run("double integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		got := slices.Map(s, func(i int) int { return i * 2 })
		assert.Equal(t, []int{2, 4, 6, 8, 10}, got)
	})

	t.Run("string to integer", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		got := slices.Map(s, func(i int) string { return strconv.Itoa(i) })
		assert.Equal(t, []string{"1", "2", "3", "4", "5"}, got)
	})

	t.Run("original type", func(t *testing.T) {
		type intSlice []int
		s := intSlice{1, 2, 3, 4, 5}
		got := slices.Map(s, func(i int) int { return i * 2 })
		assert.Equal(t, []int{2, 4, 6, 8, 10}, got)
	})
}
