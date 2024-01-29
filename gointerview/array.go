package interview

import (
	"sort"
)

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
	// dp[i] = dp[i-2]
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

func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

// H 指数
// 根据维基百科上 h 指数的定义：
// h 代表“高引用次数” ，
// 一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，
// 并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
func hIndex(citations []int) int {
	in := sort.IntSlice(citations)
	sort.Sort(in)
	h := 0
	num := len(citations)
	for i := num - 1; i >= 0; i-- {
		if citations[i] >= num-i {
			h++
		} else {
			return h
		}
	}
	return h
}

// 逻辑证明:
// 1. a>b,时，肯定能达到
// 1.
// 2.
func canCompleteCircuit(gas []int, cost []int) int {
	nums := len(gas)
	index := -1
	total := 0
	for i := 0; i < nums; i++ {
		cost[i] = gas[i] - cost[i]
		total += cost[i]
		if cost[i] > 0 {
			index = max(index, cost[i])
		}
	}
	if total >= 0 {
		return index
	}
	return -1
}

func trap(arr []int) int {
	n := len(arr)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	max := 0
	for i := 0; i < n; i++ {
		if arr[i] >= max {
			max = arr[i]
		}
		leftMax[i] = max
	}
	max = 0
	for i := n - 1; i >= 0; i-- {
		if arr[i] >= max {
			max = arr[i]
		}
		rightMax[i] = max
	}

	total := 0
	for i := 0; i < n; i++ {
		total += min(rightMax[i], leftMax[i]) - arr[i]
	}

	return total
}

func GetNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1 // 模板字符串的第一个前后缀的长度为设置为-1
	k := 0       // k为最大前缀长度
	for i := 0; i < len(pattern)-1; i++ {
		k = next[i]
		for k > -1 && pattern[i] != pattern[k] {
			k = next[k]
		}
		if k == -1 {
			next[i+1] = 0
		} else {
			next[i+1] = k + 1
		}
	}
	return next
}

func strStr(text string, pattern string) int {
	next := GetNext(pattern)
	m := len(pattern)
	for i, j := 0, 0; i < len(text); i++ {
		for j > -1 && text[i] != pattern[j] {
			j = next[j]
		}
		if j == -1 {
			j = 0
		} else {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

// 盛水最多最多的容器
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	res := 0
	for i < j {
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
		res = max(res, min(height[i], height[j])*(j-i))
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end := 0, 0
	sum := 0
	ans := n + 1
	for end < target {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == n+1 {
		ans = 0
	}
	return ans
}
