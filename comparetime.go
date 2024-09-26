package comparetime

import "time"

func lessThanSign(a time.Time, b time.Time) bool {
	return a.Before(b)
}

func lessThanOrEqualSign(a time.Time, b time.Time) bool {
	return a.Before(b) || a.Equal(b)
}

func equalSign(a time.Time, b time.Time) bool {
	return a.Equal(b)
}

func greaterThanSign(a time.Time, b time.Time) bool {
	return a.After(b)
}

func greaterThanOrEqualSign(a time.Time, b time.Time) bool {
	return a.After(b) || a.Equal(b)
}
