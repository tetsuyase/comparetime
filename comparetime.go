package comparetime

import "time"

// a < b
func lessThanSign(a time.Time, b time.Time) bool {
	return a.Before(b)
}

// a <= b
func lessThanOrEqualSign(a time.Time, b time.Time) bool {
	return a.Before(b) || a.Equal(b)
}

// a == b
func equalSign(a time.Time, b time.Time) bool {
	return a.Equal(b)
}

// a > b
func greaterThanSign(a time.Time, b time.Time) bool {
	return a.After(b)
}

// a >= b
func greaterThanOrEqualSign(a time.Time, b time.Time) bool {
	return a.After(b) || a.Equal(b)
}
