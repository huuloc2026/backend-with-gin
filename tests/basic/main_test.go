package basic

import "testing"

// Function to calculate the sum of two numbers
func TestAdD(t *testing.T) {
	var (
		input  = 5
		input2 = 5
		output = 10
	)
	if AddOne(input, input2) != output {
		t.Error("Test Failed")
	} else {
		t.Log("Test Passed")
	}
}
