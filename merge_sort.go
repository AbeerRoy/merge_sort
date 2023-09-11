package merge_sort

// Splitting the user input inputArray in to two arrays
func MergeSort(items []int64) []int64 {
	var num = int64(len(items))

	if num == 1 {
		return items
	}

	middle := int64(num / 2)
	var (
		left  = make([]int64, middle)
		right = make([]int64, num-middle)
	)
	for i := int64(0); i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

// joins the two sorted arrays to one single array
func merge(left, right []int64) []int64 {
	result := make([]int64, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
