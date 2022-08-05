package main

import "testing"

// 😎 Debugging SUCKS ！Testing ROCKS !!

// 表格驱动测试
func TestTriangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		//{12, 35, 0},
		{30000, 40000, 50000},
	}

	for _, tt := range test {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}
