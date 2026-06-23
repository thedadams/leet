impl Solution {
    pub fn remove_outer_parentheses(s: String) -> String {
        let mut openParen = 0;
        let mut result = String::with_capacity(s.len());
        for b in s.as_bytes().into_iter() {
            if *b == '(' as u8 {
                if openParen > 0 {
                    unsafe {
                        result.push(std::char::from_u32_unchecked(*b as u32));
                    }
                }
                openParen += 1;
            } else {
                openParen -= 1;
                if openParen > 0 {
                    unsafe {
                        result.push(std::char::from_u32_unchecked(*b as u32));
                    }
                }
            }
        }

        result
    }
}
