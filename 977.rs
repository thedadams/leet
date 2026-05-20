impl Solution {
    pub fn sorted_squares(nums: Vec<i32>) -> Vec<i32> {
        let mut pos = 0;
        let mut neg = 0;
        
        let mut result = Vec::with_capacity(nums.len());

        for num in nums.iter() {
            if *num < 0 {
                pos += 1;
            } else {
                break;
            }
        }

        neg = pos - 1;
        
        while result.len() < nums.len() {
            if neg < nums.len() && (pos >= nums.len() || -nums[neg] < nums[pos]) {
                result.push(nums[neg] * nums[neg]);
                neg -= 1;
            } else {
                result.push(nums[pos] * nums[pos]);
                pos += 1;
            }
        }

        result
    }
}
