impl Solution {
    pub fn add_to_array_form(num: Vec<i32>, mut k: i32) -> Vec<i32> {
        let mut result = num.clone();
        let mut carry = 0;
        let mut digit;
        let mut idx = num.len() - 1;

        while k > 0 || carry > 0 {
            digit = k % 10;
            k /= 10;

            if idx >= num.len() {
                result.splice(0..0, vec![0]);
                idx = 0;
            }

            result[idx] += digit + carry;
            carry = result[idx] / 10;
            result[idx] %= 10;
            idx -= 1;
        }

        result
    }
}
