package interview

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

// func length
func lengthOfLongestSubstring(s string) int {
<<<<<<< HEAD
	if len(s) == 1 {
		return 1
	}
	l, r := 0, 1
=======
	num := len(s)
	if num <= 1 {
		return len(s)
	}
	l, r := 0, 0
>>>>>>> 3b539d9232bf7d069bd1cbbfc11db94da6beec79
	res := 0
	m := make(map[byte]int)
	for r < len(s) {
		index, ok := m[s[r]]
<<<<<<< HEAD
		if !ok {
			m[s[r]] = r
		} else {
=======
		if ok && index >= l {
>>>>>>> 3b539d9232bf7d069bd1cbbfc11db94da6beec79
			res = max(res, r-l)
			l = index + 1
		}
		m[s[r]] = r
		r++
	}
	res = max(res, r-l)
	return res
}
