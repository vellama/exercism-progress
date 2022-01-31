package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	_, isInRows := cb[rank]
	var sum int = 0

	if !isInRows {
		return sum
	}

	for i := 0; i < len(cb[rank]); i++ {
		if cb[rank][i] {
			sum++
		}
	}

	return sum
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var sum int = 0

	if file < 1 || file > 8 {
		return sum
	}

	for _, rank := range cb {
		if rank[file-1] {
			sum++
		}
	}

	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var sum int = 0

	for _, rank := range cb {
		sum += len(rank)
	}

	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var sum int = 0

	for key := range cb {
		sum += CountInRank(cb, key)
	}

	return sum
}
