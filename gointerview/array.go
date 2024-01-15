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
	if len(nums) <= 2 {
		return len(nums)
	}

	resultIdx := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[resultIdx-2] {
			nums[resultIdx] = nums[i]
			resultIdx++
		}
	}
	return resultIdx
}

func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	count := len(nums)
	k = k % count
	s := make([]int, k)
	index := count - k
	copy(s, nums[index:])
	copy(nums[k:], nums[0:index])
	copy(nums[:k], s)
}
