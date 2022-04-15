package slices

func All[E any](s []E, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[E any](s []E, f func(E) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if vs == v {
			return i
		}
	}
	return -1
}

func Include[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}

func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	tot := len(s) + len(v)
	if tot <= cap(s) {
		s2 := s[:tot]
		copy(s2[i+len(v):], s[i:])
		copy(s2[i:], v)
		return s2
	}
	s2 := make(S, tot)
	copy(s2, s[:i])
	copy(s2[i:], v)
	copy(s2[i+len(v):], s[i:])
	return s2
}

func Delete[S ~[]E, E any](s S, i, j int) S {
	return append(s[:i], s[j:]...)
}

func Clone[S ~[]E, E any](s S) S {
	if s == nil {
		return nil
	}
	return append(S([]E{}), s...)
}

func Map[S1 ~[]E1, S2 []E2, E1, E2 any](s S1, f func(E1) E2) S2 {
	dst := make(S2, len(s), len(s))
	for i, v := range s {
		dst[i] = f(v)
	}
	return dst
}

func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	dst := make(S, 0)
	for _, v := range s {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return dst
}
