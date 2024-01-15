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
