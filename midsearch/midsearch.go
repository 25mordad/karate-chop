package midsearch

//Chop return int
func Chop(item int, items []int) int {
	min := 0
	max := len(items)
	mid := (min + max) / 2
	// fmt.Println("#items", min, mid, max)

	for min < max {
		// fmt.Println("#items", min, mid, max)

		if items[mid] == item {

			return mid
		}

		if items[mid] > item {
			max = mid - 1
		} else {
			min = mid + 1
		}
		mid = (min + max) / 2
	}

	if max == len(items) || items[max] != item {
		return -1
	}

	return max
}
