impl Solution {
    pub fn common_chars(words: Vec<String>) -> Vec<String> {
        let word_count = words.len() as i32;
        let mut letters: std::collections::HashMap<i32, i32> = std::collections::HashMap::with_capacity(26 * 5);
        for w in words.into_iter() {
            let mut letters_this_word: std::collections::HashMap<i32, i32> = std::collections::HashMap::with_capacity(26 * 5);
            for b in w.as_bytes().into_iter() {
                let b = *b as i32;
                let i = match letters_this_word.get(&b) {
                    Some(i) => *i,
                    None => 0,
                };

                let idx = b + 97 * i;
                let c = match letters.get(&idx) {
                    Some(c) => *c,
                    None => 0,
                };
                letters.insert(idx, c+1);
                letters_this_word.insert(b, i+1);
            }
        }

        let mut result = Vec::new();
        for (k, v) in letters.into_iter() {
            if v == word_count {
                result.push(String::from_utf8(vec![((k % 97) + 97) as u8]).unwrap());
            }
        }

        result
    }
}
