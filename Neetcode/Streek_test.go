package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "simpleTest",
			nums: []int{1, 2, 3, 4, 1, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums...)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStreek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestConsecutive(1, 2, 3, 4, 1, 2)
	}
}
