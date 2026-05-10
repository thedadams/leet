impl Solution {
    pub fn min_deletion_size(strs: Vec<String>) -> i32 {
        let mut result = 0;

        // unsafe jus to make it a little faster
        unsafe {
            for i in 0..strs[0].len() {
                for j in 0..(strs.len()-1) {
                    if strs[j].get_unchecked(i..=i) > strs[j+1].get_unchecked(i..=i) {
                        result += 1;
                        break;
                    }
                }
            }
        }

        result
    }
}
