package lab2

// Count returns how many times value a is included in the slice s.
func Count(s []int, a int) int {
	r := 0
	for _, e := range s {
		if a == e {
			r++
		}
	}
	return r
}
