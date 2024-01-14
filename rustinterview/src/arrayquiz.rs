/*!
  # ArrayQuiz Module

  `arrayquiz` 是一个提供数组操作函数的模块，例如合并两个排序好的数组。
*/


    /// 合并两个已排序的数组。
    /// 这个函数实现了 LeetCode 上的 "合并两个有序数组" 问题。
    /// 详细描述和示例可以在这里找到：[LeetCode 题目](https://leetcode.cn/problems/merge-sorted-array).
    /// # 参数
    ///
    /// * `nums1` - 第一个排序数组的可变引用，其中含有足够的空间来容纳两个数组合并后的所有元素。
    /// * `m` - `nums1` 中已排序元素的数量。
    /// * `nums2` - 第二个排序数组的可变引用。
    /// * `n` - `nums2` 中元素的数量。
    ///
    /// # 示例
    ///
    /// ```
    /// use rustinterview::arrayquiz::merge_sorted_array; 
    /// let mut nums1 = vec![1,2,3,0,0,0];
    /// let mut nums2 = vec![2,5,6];
    /// merge_sorted_array(&mut nums1, 3, &mut nums2, 3);
    /// assert_eq!(nums1, vec![1,2,2,3,5,6]);
    /// ```
pub fn merge_sorted_array(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
     let mut p1 = m as usize - 1;
     let mut p2 = n as usize - 1;
     for i in (0..(m as usize + n as usize)).rev() {
        if p2 == usize::MAX {
            break;
        }
        if nums1[p1]>=nums2[p2] {
            nums1[i] = nums1[p1];
            p1 = p1.wrapping_sub(1);
        }else{
            nums1[i] = nums2[p2];
            p2 = p2.wrapping_sub(1);
        }
    }
}


#[cfg(test)]
mod tests {
    use std::vec;

    #[test]
    fn test_merge_sorted_array() {
        let mut nums1 = vec![1,2,3,0,0,0];
        let m = 3;
        let mut nums2 = vec![2,5,6];
        let n = 3;
        super::merge_sorted_array(& mut nums1, m, & mut nums2, n);
        assert_eq!(nums1, vec![1,2,2,3,5,6]); 
    }
}