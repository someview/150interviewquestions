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
