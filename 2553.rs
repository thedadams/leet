impl Solution {
    pub fn separate_digits(nums: Vec<i32>) -> Vec<i32> {
        let mut result = Vec::with_capacity(nums.len());
        let mut single = Vec::new();

        for mut n in nums {
            while n > 0 {
                single.push(n % 10);
                n = n / 10;
            }

            for n in single.iter().rev() {
                result.push(*n);
            }

            single.clear();
        }

        result
    }
}
