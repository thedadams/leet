impl Solution {
    pub fn is_alien_sorted(mut words: Vec<String>, order: String) -> bool {
        let mut alphabet: std::collections::HashMap<char,_> = std::collections::HashMap::with_capacity(order.len());
        let mut idx = 0;
        for c in order.chars() {
            alphabet.insert(c, idx);
            idx += 1;
        }

        let mut new_words = Vec::with_capacity(words.len());
        for s in words.iter() {
            let mut new_s = String::with_capacity(s.len());
            for c in s.chars() {
                if let Some(h) = alphabet.get(&c) {
                    unsafe {
                        new_s.push(char::from_u32_unchecked(97 + *h));
                    }
                }
            }

            new_words.push(new_s);
        }

        new_words.is_sorted()
    }
}
