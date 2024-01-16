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

// 最优解, 股票价格相当于一系列的折线图,寻找递增子序列中的最大值，只需要遍历一次就足够了
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			maxProfit = max(maxProfit, price-minPrice)
		}
	}
	return maxProfit
}

// 买卖股票的最佳时机 II
func accMaxProfit(prices []int) int {
	days := len(prices)
	if days <= 1 {
		return 0
	}
	totalProfit := 0
	for day := 0; day < days-1; day++ {
		if prices[day] < prices[day+1] {
			totalProfit += prices[day+1] - prices[day]
		}
	}

	return totalProfit
}

// 跳跃游戏
func canJump(nums []int) bool {
	// 逆序，寻找为0的位置，前面的元素是否可以跳过为0的位置
	// 动态规划
	// dp[i] = max(dp[i-1]+ nums[i]+i)
	// dp[i-1] = max(dp[i-2]+nums[i-1]+1)
	// dp[0] = nums[0]
	count := len(nums)
	if count == 1 {
		return true
	}
	dp := 0
	for i := 0; i < count-1; i++ {
		dp = max(dp, nums[i]+i)
		if dp == i {
			return false
		}
		if dp >= count-1 {
			return true
		}
	}
	return false
}
