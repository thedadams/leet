impl Solution {
    pub fn sort_array_by_parity(mut nums: Vec<i32>) -> Vec<i32> {
        let mut idx = 0;
        let mut end = nums.len() - 1;
        while idx < end {
            if nums[idx] % 2 == 1 {
                nums.swap(idx, end);
                end -= 1;
            } else {
                idx += 1;
            }
        }
        
        nums
    }
}
