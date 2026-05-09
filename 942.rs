impl Solution {
    pub fn di_string_match(mut s: String) -> Vec<i32> {
        let mut result: Vec<i32> = (0..(s.len()+1) as i32).collect();

        let mut i = 0;
        while let Some(c) = s.get(i..i+1) {
            if c == "I" {
                i += 1;
                continue;
            }

            let mut j = i + 1;
            while let Some(d) = s.get(j..j+1) {
                if d == "I" {
                    break;
                }
                j += 1;
            }

            j += 1;
            for n in (i..j).rev() {
                result[i] = n as i32;
                i += 1;
            }
        }
        
        result
    }
}
