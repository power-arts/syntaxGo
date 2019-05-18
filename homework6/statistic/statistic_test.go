package statistic

/*
 * Syntax Go Homework 6
 * Stepanov Anton, 18.05.2019
 * 1 пункт
 */

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

func TestSum(t *testing.T) {
	var v float64
	v = Sum([]float64{1.5, 2, 2.5, 4, 5})
	if v != 15 {
		t.Error("Expected 15, got ", v)
	}
}
