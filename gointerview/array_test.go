package interview

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergemergeSortedArray(t *testing.T) {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	mergeSortedArray(nums1, m, nums2, n)
}

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	assert.Equal(t, 5, removeElement(nums, val))
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	assert.Equal(t, 5, removeDuplicates(nums))
	assert.ElementsMatch(t, nums[:5], []int{0, 1, 2, 3, 4})
}

func TestRemoveLimitedDuplicated(t *testing.T) {
	nums := []int{0, 0, 0, 1, 1, 1, 1, 2, 3, 3}
	assert.Equal(t, 7, removeLimitedDuplicates(nums))
	fmt.Println("nums:", nums)
	assert.ElementsMatch(t, nums[:7], []int{0, 0, 1, 1, 2, 3, 3})
	nums = []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, 5, removeLimitedDuplicates(nums))
	assert.ElementsMatch(t, nums[:5], []int{1, 1, 2, 2, 3})
}

func TestGetMajorElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	assert.Equal(t, 2, majorityElement(nums))
}

func TestMaxProfit(t *testing.T) {
	prices := []int{1, 2}
	assert.Equal(t, 1, maxProfit(prices))
	prices = []int{7, 6, 5}
	assert.Equal(t, 0, maxProfit(prices))
	prices = []int{7, 1, 3, 2, 5, 4}
	assert.Equal(t, 4, maxProfit(prices))
}

func TestAccMaxProfit(t *testing.T) {
	prices := []int{1, 2}
	assert.Equal(t, 1, accMaxProfit(prices))
	prices = []int{7, 6, 5}
	assert.Equal(t, 0, accMaxProfit(prices))
	prices = []int{7, 1, 3, 2, 5, 4}
	assert.Equal(t, 5, accMaxProfit(prices))
}

func TestCanJump(t *testing.T) {
	grids := []int{3, 2, 1, 0, 4}
	assert.False(t, canJump(grids))
}

func TestHIndex(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	assert.Equal(t, 3, hIndex(citations))
}

func TestCanRunCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))

	gas = []int{2, 3, 4}
	cost = []int{3, 4, 3}
	assert.Equal(t, -1, canCompleteCircuit(gas, cost))

	gas = []int{1, 4, 4, 3}
	cost = []int{3, 4, 1, 2}
	assert.Equal(t, 3, canCompleteCircuit(gas, cost))
}

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	assert.Equal(t, 6, trap(height))
	height = []int{4, 2, 0, 3, 2, 5}
	assert.Equal(t, 9, trap(height))
}

func Test_strStr(t *testing.T) {
	text := "mississippi"
	pattern := "issip"
	assert.Equal(t, 4, strStr(text, pattern))
}

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))
}
