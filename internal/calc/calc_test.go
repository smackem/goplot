package calc

import "testing"

func TestCalculator_Evaluate(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		c    *Calculator
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Evaluate(tt.args.src)
		})
	}
}
