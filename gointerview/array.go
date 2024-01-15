package interview

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	for index := m + n - 1; index >= 0; index-- {
		if p2 < 0 {
			return
		}
		if p1 < 0 {
			nums1[index] = nums2[p2]
			p2--
		} else if nums1[p1] >= nums2[p2] {
			nums1[index] = nums1[p1]
			p1--
		} else {
			nums1[index] = nums2[p2]
			p2--
		}
	}
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}
	return left
}

// 移除数组中的重复元素
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return l
	}
	resultIdx := 1
	for i := 1; i < l; i++ {
		if nums[i-1] != nums[i] {
			nums[resultIdx] = nums[i]
			resultIdx++
		}
	}
	return resultIdx
}

// 删除数组中的重复元素，使其最多出现两次
func removeLimitedDuplicates(nums []int) int {
	// GPT给出的最优解法
	// if len(nums) <= 2 {
	//     return len(nums)
	// }

	// resultIdx := 2

	// for i := 2; i < len(nums); i++ {
	//     if nums[i] != nums[resultIdx-2] {
	//         nums[resultIdx] = nums[i]
	//         resultIdx++
	//     }
	// }

	// return resultIdx
	count := len(nums)
	if count <= 1 {
		return count
	}
	resultIdx := 1
	startIdx := 0
	for nowIndex := 1; nowIndex < count; nowIndex++ {
		if nums[nowIndex] != nums[nowIndex-1] {
			nums[resultIdx] = nums[nowIndex]
			resultIdx++
			startIdx = nowIndex
			continue
		}
		if startIdx+1 >= nowIndex {
			nums[resultIdx] = nums[nowIndex]
			resultIdx++
		}
	}
	return resultIdx
}

func GetMajorElement(nums []int) int {
	// 假设大小为n-1的nums的最多元素为m,那么，nums[n-1] + i, 其最大元素一定为m，或者max(nums[n-1])
	if len(nums) == 1 {
		return 0
	}
	end := len(nums) - 1
	mid := end / 2
	privot := nums[mid]
	for i := 0; i < end; i++ {
		if nums[i] >= privot {
			nums[privot+1] = nums[mid]
			mid++
		}
	}
	return nums[len(nums)/2]
}
