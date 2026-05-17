impl Solution {
    pub fn repeated_n_times(nums: Vec<i32>) -> i32 {
        let n = nums.len() / 3;

        for i in 0..n {
            if nums[3*i] == nums[3*i+1] || nums[3*i] == nums[3*i+2] {
                return nums[3*i];
            }
            if nums[3*i+1] == nums[3*i+2] {
                return nums[3*i+1];
            }
        }

        nums[nums.len()-1]
    }
}
