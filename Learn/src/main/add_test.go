package main

import "testing"

//表格驱动测试
func TestAdd(t *testing.T) {
	tests := []struct{a, b, c int} {
		{1,2,3},
	}

	for _, tt := range tests {
		if res := tt.a + tt.b; res != tt.c {
			t.Errorf("expected: %d; but actual: %d\n", tt.c, res)
		}
	}
}