package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	boxes := make([]map[byte]struct{}, 9)

	for i := range 9 {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		boxes[i] = make(map[byte]struct{})
	}

	for i := range 9 {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			boxIdx := (i/3)*3 + (j / 3)

			if _, ok := rows[i][val]; ok {
				return false
			}
			if _, ok := cols[j][val]; ok {
				return false
			}
			if _, ok := boxes[boxIdx][val]; ok {
				return false
			}
			rows[i][val] = struct{}{}
			cols[j][val] = struct{}{}
			boxes[boxIdx][val] = struct{}{}
		}
	}
	return true
}

// ВИЗУАЛЬНОЕ ПРЕДСТАВЛЕНИЕ 3x3 box индексов:
//
// Индексы боксов (boxIdx) для каждой клетки:
//
// [ 0 | 0 | 0 | 1 | 1 | 1 | 2 | 2 | 2 ]
// [ 0 | 0 | 0 | 1 | 1 | 1 | 2 | 2 | 2 ]
// [ 0 | 0 | 0 | 1 | 1 | 1 | 2 | 2 | 2 ]
// [ 3 | 3 | 3 | 4 | 4 | 4 | 5 | 5 | 5 ]
// [ 3 | 3 | 3 | 4 | 4 | 4 | 5 | 5 | 5 ]
// [ 3 | 3 | 3 | 4 | 4 | 4 | 5 | 5 | 5 ]
// [ 6 | 6 | 6 | 7 | 7 | 7 | 8 | 8 | 8 ]
// [ 6 | 6 | 6 | 7 | 7 | 7 | 8 | 8 | 8 ]
// [ 6 | 6 | 6 | 7 | 7 | 7 | 8 | 8 | 8 ]
//
// boxIdx для клетки [i][j] вычисляется так: (i/3)*3 + (j/3)

func Sudoku(board [][]string) bool {

	rows := make([]map[string]bool, 9)
	coals := make([]map[string]bool, 9)
	boxes := make([]map[string]bool, 9)

	for i := range board {
		rows[i] = make(map[string]bool)
		coals[i] = make(map[string]bool)
		boxes[i] = make(map[string]bool)
	}

	for i := range rows {
		for j := range coals {
			num := board[i][j]

			if num == "." {
				continue
			}

			boxIndex := (i/3)*3 + (j / 3)

			if rows[i][num] || coals[j][num] || boxes[boxIndex][num] {
				return false
			}
			rows[i][num] = true
			coals[j][num] = true
			rows[boxIndex][num] = true

		}
	}
	return true
}
