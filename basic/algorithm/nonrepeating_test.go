package main

import (
	"testing"
)

func TestSubStr(t *testing.T) {
	test := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是BIOS,收到请回答～", 14},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range test {
		actual := longestSubstring(tt.s)
		if tt.ans != actual {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {

	s := "什么文具最谨慎？尺子，因为它知道分寸"
	// 加长待测试的字符串
	for i := 0; i < 15; i++ {
		s = s + s
	}
	ans := 18

	b.Logf("len(s) = %d", len(s))
	// 重制时间，只计算下面真正执行运算的部分
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := longestSubstring(s)
		if actual != ans {
			b.Errorf("got %d for input something; expected %d", actual, ans)
		}
	}
}
