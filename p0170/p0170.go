package p0170

func reversePairs(record []int) int {
	if len(record) < 2 {
		return 0
	}
	temp := make([]int, len(record))
	return mergeSort(record, 0, len(record)-1, temp)
}

func mergeSort(record []int, left, right int, temp []int) int {
	// left, right是从0起始的数组下标, 这里temp用来存放归并排序后的数组，只是常规做法，实际可以不用temp数组的

	// 拆分到单一元素时，无逆序对
	if left == right {
		return 0
	}

	mid := (left + right) / 2
	leftCount := mergeSort(record, left, mid, temp)
	rightCount := mergeSort(record, mid+1, right, temp)
	reverseCount := leftCount + rightCount
	if record[mid] <= record[mid+1] {
		return reverseCount
	}
	return reverseCount + countDuringMerge(record, left, mid, right, temp)
}

func countDuringMerge(record []int, left, mid, right int, temp []int) int {
	reverseCount := 0
	for i := left; i <= right; i++ {
		temp[i] = record[i]
	}

	// 双指针，指向两个有序数组起点
	i, j := left, mid+1
	for k := left; k <= right; k++ {
		// 左数组已归并完比，直接归并右数组即可
		if i > mid {
			record[k] = temp[j]
			j++
		} else if j > right { // 右边数组归并完毕，则直接归并左数组即可
			record[k] = temp[i]
			i++
		} else if temp[i] <= temp[j] {
			record[k] = temp[i]
			i++
		} else {
			record[k] = temp[j]
			j++
			reverseCount += mid - i + 1
		}
	}
	return reverseCount
}
