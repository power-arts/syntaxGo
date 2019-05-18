package main

/*
 * Syntax Go Homework 6
 * Stepanov Anton, 18.05.2019
 * 4 пункт
 */

import "testing"

func TestQuadraticEquations(t *testing.T) {
	var x1 float64
	var x2 float64
	x1, x2, _ = quadraticEquations(4, -7, -2)
	if x1 != 2 || x2 != -0.25 {
		t.Errorf("X1 must be 2 got %v, X2 must be -0.25 got %v", x1, x2)
	}
}
