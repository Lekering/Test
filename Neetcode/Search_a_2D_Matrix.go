package main

func Binary2dMatrix(array [][]int, target int) bool {
	rows, cols := len(array), len(array[0])
	l, r := 0, rows*cols-1
	for l <= r {
		m := l + (r-l)/2
		row, col := m/cols, m%cols
		if array[row][col] == target {
			return true
		} else if array[row][col] > target {
			r = m - 1
		} else if array[row][col] < target {
			l = m + 1
		}
	}
	return false
}
