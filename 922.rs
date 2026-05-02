impl Solution {
    pub fn sort_array_by_parity_ii(mut nums: Vec<i32>) -> Vec<i32> {
        let mut i = 0;
        let mut j = 1;

        while i < nums.len() {
            if nums[i] % 2 != (i % 2) as i32 {
                while nums[j] % 2 == nums[i] % 2 {
                    j += 2;
                }

                nums.swap(i, j);
            }

            i += 2;
        }

        nums
    }
}
