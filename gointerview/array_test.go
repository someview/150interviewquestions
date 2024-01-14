package interview

import (
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
