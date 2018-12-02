package validSudoku

// Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

// The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

// Example 1:

// Input:
// [
//   ["5","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: true
// Example 2:

// Input:
// [
//   ["8","3",".",".","7",".",".",".","."],
//   ["6",".",".","1","9","5",".",".","."],
//   [".","9","8",".",".",".",".","6","."],
//   ["8",".",".",".","6",".",".",".","3"],
//   ["4",".",".","8",".","3",".",".","1"],
//   ["7",".",".",".","2",".",".",".","6"],
//   [".","6",".",".",".",".","2","8","."],
//   [".",".",".","4","1","9",".",".","5"],
//   [".",".",".",".","8",".",".","7","9"]
// ]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being 
//     modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

func isValidSudoku(board [][]byte) bool {
    for i := 0; i < 9; i++ {
        if !validRow(board, i) || !validColumn(board, i) {
            return false
        }
    }
    for i := 0; i < 9; i+=3 {
        for j := 0; j < 9; j+=3 {
            if !validSub(board, i, j) {
                return false
            }
        }
    }
    return true
}

func validRow(board [][]byte, row int) bool {
    set := make(map[byte]struct{}, 0)
    for _, i := range board[row] {
        if i >= '1' && i <= '9' {
            if _, ok := set[i]; ok {
                return false
            } else {
                set[i] = struct{}{}
            }
        }
    }
    return true
}

func validColumn(board [][]byte, col int) bool {
    set := make(map[byte]struct{}, 0)
    for _, row := range board {
        i := row[col]
        if i >= '1' && i <= '9' {
            if _, ok := set[i]; ok {
                return false
            } else {
                set[i] = struct{}{}
            }
        }
    }
    return true
}

func validSub(board [][]byte, row, col int) bool {
    set := make(map[byte]struct{}, 0)
    for m := 0; m < 3; m++ {
        for n := 0; n < 3; n++ {
            i := board[row+m][col+n]
            if i >= '1' && i <= '9' {
                if _, ok := set[i]; ok {
                    return false
                } else {
                    set[i] = struct{}{}
                }
            }
        }
    }
    return true
}