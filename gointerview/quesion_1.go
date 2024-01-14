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
