package main

import "testing"

func TestComputeData(t *testing.T) {
	result := ComputeData(2, 3)
	if result != 5 {
		t.Errorf("Expected %d, got %d", 2*3, result)
	}
}
