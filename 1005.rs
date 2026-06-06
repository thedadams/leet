impl Solution {
    pub fn largest_sum_after_k_negations(mut nums: Vec<i32>, mut k: i32) -> i32 {
        nums.sort();

        let mut result = 0;
        let mut lastNeg = 0;

        for m in nums.iter() {
            let n = *m;
            if n < 0 {
                if k > 0 {
                    result -= n;
                    k -= 1;
                } else {
                    result += n;
                }
                lastNeg = -n;
            } else if n == 0 {
                k = 0;
            } else {
                if k % 2 == 1 {
                    if lastNeg > 0 && lastNeg < n {
                        result -= 2 * lastNeg;
                        result += n;
                    } else {
                        result -= n;
                    }
                    k = 0;
                } else {
                    result += n;
                }
            }
        }

        if k % 2 == 1 {
            // The only way to get here is to have more indicies to flip, but everything was negative.
            // So flip the largest negative.
            result += 2 * nums[nums.len() - 1];
        }

        result
    }
}
