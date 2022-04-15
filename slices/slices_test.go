package slices_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"

	"johnmanjiro13/gocol/slices"
)

func isEven(i int) bool {
	return i%2 == 0
}

func TestAll(t *testing.T) {
	tests := map[string]struct {
		s    []int
		f    func(int) bool
		want bool
	}{
		"nil slice":    {nil, isEven, true},
		"blank slice":  {nil, isEven, true},
		"all even":     {[]int{2, 4, 6, 8}, isEven, true},
		"contains odd": {[]int{1, 2, 4, 6, 8}, isEven, false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, slices.All(tt.s, tt.f))
		})
	}
}

func TestAny(t *testing.T) {
	tests := map[string]struct {
		s    []int
		f    func(int) bool
		want bool
	}{
		"nil slice":    {nil, isEven, false},
		"blank slice":  {nil, isEven, false},
		"all even":     {[]int{2, 4, 6, 8}, isEven, true},
		"contains odd": {[]int{1, 2, 4, 6, 8}, isEven, true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, slices.Any(tt.s, tt.f))
		})
	}
}

var indexTests = map[string]struct {
	s    []int
	v    int
	want int
}{
	"nil slice":        {nil, 0, -1},
	"blank slice":      {[]int{}, 0, -1},
	"single value":     {[]int{1, 2, 3}, 2, 1},
	"duplicate value":  {[]int{1, 2, 2, 3}, 2, 1},
	"duplicate value2": {[]int{1, 2, 3, 2}, 2, 1},
	"not existed":      {[]int{1, 2, 3}, 4, -1},
}

func TestIndex(t *testing.T) {
	for name, tt := range indexTests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, slices.Index(tt.s, tt.v))
		})
	}
}

func TestInclude(t *testing.T) {
	for name, tt := range indexTests {
		t.Run(name, func(t *testing.T) {
			got := slices.Include(tt.s, tt.v)
			assert.Equal(t, tt.want != -1, got)
		})
	}
}

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

func TestFilter(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		got := slices.Filter(s, isEven)
		assert.Equal(t, []int{2, 4}, got)
	})

	t.Run("string", func(t *testing.T) {
		s := []string{"apple", "banana", "avocado"}
		got := slices.Filter(s, func(s string) bool { return strings.HasPrefix(s, "a") })
		assert.Equal(t, []string{"apple", "avocado"}, got)
	})

	t.Run("original type", func(t *testing.T) {
		type intSlice []int
		s := intSlice{1, 2, 3, 4, 5}
		got := slices.Filter(s, isEven)
		assert.Equal(t, intSlice{2, 4}, got)
	})
}
