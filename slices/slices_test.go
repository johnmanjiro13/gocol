package slices_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestInsert(t *testing.T) {
	tests := map[string]struct {
		s    []int
		i    int
		add  []int
		want []int
	}{
		"first": {
			[]int{1, 2, 3, 4},
			0,
			[]int{5, 6},
			[]int{5, 6, 1, 2, 3, 4},
		},
		"middle": {
			[]int{1, 2, 3, 4},
			2,
			[]int{5, 6},
			[]int{1, 2, 5, 6, 3, 4},
		},
		"last": {
			[]int{1, 2, 3, 4},
			4,
			[]int{5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := slices.Insert(tt.s, tt.i, tt.add...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := map[string]struct {
		s    []int
		i    int
		j    int
		want []int
	}{
		"first": {
			[]int{1, 2, 3},
			0,
			0,
			[]int{1, 2, 3},
		},
		"only one": {
			[]int{1, 2, 3, 4, 5},
			0,
			1,
			[]int{2, 3, 4, 5},
		},
		"middle": {
			[]int{1, 2, 3, 4, 5},
			2,
			3,
			[]int{1, 2, 4, 5},
		},
		"last": {
			[]int{1, 2, 3, 4, 5},
			4,
			4,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := slices.Delete(tt.s, tt.i, tt.j)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClone(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.Clone(s1)
	assert.Equal(t, s1, s2)

	s1[0] = 6
	assert.NotEqual(t, s1, s2)

	assert.Nil(t, slices.Clone([]int(nil)))

	assert.NotNil(t, slices.Clone(s1[:0]))
	assert.Equal(t, 0, len(slices.Clone(s1[:0])))
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
