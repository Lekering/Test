package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		nums []int
		want int
	}{
		{
			name: "simple consecutive sequence",
			nums: []int{2, 3, 4, 5},
			want: 4,
		},
		{
			name: "disjoint sequences",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{10},
			want: 1,
		},
		{
			name: "empty input",
			nums: []int{},
			want: 0,
		},
		{
			name: "duplicate numbers",
			nums: []int{1, 2, 2, 3},
			want: 3,
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
	for b.Loop() {
		longestConsecutive(1, 2, 3, 4, 1, 2)
	}
}
