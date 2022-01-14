package cards

func indexInBound(slice []int, index int) bool {
	return index > -1 && index < len(slice)
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	var inBound bool = indexInBound(slice, index)
	if !inBound {
		return 0, inBound
	}
	return slice[index], inBound
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if !indexInBound(slice, index) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var slice []int

	if length <= 0 {
		return slice
	}

	for i := 0; i < length; i++ {
		slice = append(slice, value)
	}

	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if !indexInBound(slice, index) {
		return slice
	}

	var sliceLength int = len(slice)
	var newSliceLength int = sliceLength - 1

	for i := index; i < sliceLength-1; i++ {
		slice[i] = slice[i+1]
	}

	return slice[:newSliceLength]
}
