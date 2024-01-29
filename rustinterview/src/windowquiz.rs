/*!
# windowquiz

`windowquiz` 滑动窗口算法题
 */

/** `min_sub_array_len` 求指定数组连续元素指和能达到target的最小元素的数量

# 示例
```
let target = 1;
let input = vec![1,2,3];
let answer = rustinterview::windowquiz::min_sub_array_len(target, input);

assert_eq!(1, answer);
```

*/

pub fn min_sub_array_len(target: i32, nums: Vec<i32>) -> i32 {
   let (mut start, mut sum, mut min_len) = (0, 0, i32::MAX);
   for end in 0..nums.len() {
      sum += nums[end];
      while sum >= target {
         min_len = min_len.min((end - start + 1) as i32);
         sum -= nums[start];
         start += 1;
      }
   }
   if min_len == i32::MAX {
      0
   }else{
      min_len
   }
}