package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	file_values, e := cb[file]
	counter := 0
	if e {
		for _, value := range file_values {

			if value {
				counter++
			}
		}
	}
	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	counter := 0
	for _, files := range cb {
		for index, val := range files {
			if index == rank-1 {
				if val {
					counter++
				}
			}
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	counter := 0
	for _, file := range cb {
		counter += len(file)
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, file := range cb {
		for _, v := range file {
			if v {
				counter++
			}
		}
	}
	return counter
}
