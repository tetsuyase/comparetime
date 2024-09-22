package comparetime

import "time"

func lessThanSign(left time.Time, right time.Time) bool {
	return left.Before(right)
}

func lessThanOrEqualSign(left time.Time, right time.Time) bool {
	return left.Before(right) || left.Equal(right)
}

func equalSign(left time.Time, right time.Time) bool {
	return left.Equal(right)
}

func greaterThanSign(left time.Time, right time.Time) bool {
	return left.After(right)
}

func greaterThanOrEqualSign(left time.Time, right time.Time) bool {
	return left.After(right) || left.Equal(right)
}
