package slices

func Map[S1 ~[]E1, S2 []E2, E1, E2 any](s S1, f func(E1) E2) S2 {
	dst := make(S2, len(s), len(s))
	for i, v := range s {
		dst[i] = f(v)
	}
	return dst
}
