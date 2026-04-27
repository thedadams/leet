impl Solution {
    pub fn smallest_range_i(nums: Vec<i32>, k: i32) -> i32 {
        let mut min = 10000;
        let mut max = 0;

        for n in nums.into_iter() {
            if min > n {
                min = n
            }
            if max < n {
                max = n
            }
        }

        let result = max - min - 2 * k;
        if result < 0 {
            return 0;
        }
        result
    }
}
