package test

import "testing"

func TestTest(t *testing.T) {
	a := 1
	b := 2

	expected_sum := 3
	actual_sum := a + b

	if (actual_sum != expected_sum) {
		t.Errorf("Expected sum of %d, got %d instead", expected_sum, actual_sum)
	}
}
