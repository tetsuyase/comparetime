package comparetime

import (
	"testing"
	"time"
)

func TestLessThanSign(t *testing.T) {
	t1 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00
	t2 := time.Date(2023, 9, 2, 12, 0, 0, 0, time.UTC) // 2023/09/02 12:00

	if !lessThanSign(t1, t2) {
		t.Errorf("Expected %v to be before %v", t1, t2)
	}

	if lessThanSign(t2, t1) {
		t.Errorf("Expected %v to not be before %v", t2, t1)
	}

	t3 := t1
	if lessThanSign(t1, t3) {
		t.Errorf("Expected %v to not be before %v since they are equal", t1, t3)
	}
}

func TestLessThanOrEqualSign(t *testing.T) {
	t1 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00
	t2 := time.Date(2023, 9, 2, 12, 0, 0, 0, time.UTC) // 2023/09/02 12:00

	if !lessThanOrEqualSign(t1, t2) {
		t.Errorf("Expected %v to be less than or equal to %v", t1, t2)
	}

	if lessThanOrEqualSign(t2, t1) {
		t.Errorf("Expected %v to not be less than or equal to %v", t2, t1)
	}

	t3 := t1
	if !lessThanOrEqualSign(t1, t3) {
		t.Errorf("Expected %v to be less than or equal to %v since they are equal", t1, t3)
	}
}

func TestEqualSign(t *testing.T) {
	t1 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00
	t2 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00

	if !equalSign(t1, t2) {
		t.Errorf("Expected %v to be equal to %v", t1, t2)
	}

	t3 := time.Date(2023, 9, 1, 12, 0, 0, 1, time.UTC)
	if equalSign(t1, t3) {
		t.Errorf("Expected %v to not be equal to %v", t1, t3)
	}
}

func TestGreaterThanSign(t *testing.T) {
	t1 := time.Date(2023, 9, 2, 12, 0, 0, 0, time.UTC) // 2023/09/02 12:00
	t2 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00

	if !greaterThanSign(t1, t2) {
		t.Errorf("Expected %v to be after %v", t1, t2)
	}

	if greaterThanSign(t2, t1) {
		t.Errorf("Expected %v to not be after %v", t2, t1)
	}

	t3 := t1
	if greaterThanSign(t1, t3) {
		t.Errorf("Expected %v to not be after %v since they are equal", t1, t3)
	}
}

func TestGreaterThanOrEqualSign(t *testing.T) {
	t1 := time.Date(2023, 9, 1, 12, 0, 0, 0, time.UTC) // 2023/09/01 12:00
	t2 := time.Date(2023, 9, 2, 12, 0, 0, 0, time.UTC) // 2023/09/02 12:00

	if greaterThanOrEqualSign(t1, t2) {
		t.Errorf("Expected %v to be greater than or equal to %v", t1, t2)
	}

	t3 := t1
	if !greaterThanOrEqualSign(t1, t3) {
		t.Errorf("Expected %v to be equal to %v", t1, t3)
	}
	if !greaterThanOrEqualSign(t2, t1) {
		t.Errorf("Expected %v to not be greater than or equal to %v", t2, t1)
	}
}
