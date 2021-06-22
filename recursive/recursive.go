package recursive

func Chop(item int, items []int) int {

	min := 0
	max := len(items)
	mid := (min + max) / 2
	// fmt.Println("#items", min, mid, max)
	// fmt.Println("#items", items)

	if min < max {

		if items[mid] == item {
			return mid
		} else {
			if items[mid] > item {
				return Chop(item, items[:mid])

			} else if items[mid] < item {
				return mid + Chop(item, items[mid:])
			}
		}

	}
	return -1

}
