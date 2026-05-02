impl Solution {
    pub fn reverse_only_letters(mut s: String) -> String {
        unsafe {
            let mut s = s.as_bytes_mut();
    
            let mut start = 0;
            let mut end = s.len() - 1;
            while start < end {
                if s[start] >= b'A' && s[start] <= b'Z' || s[start] >= b'a' && s[start] <= b'z' {
                    while !(s[end] >= b'A' && s[end] <= b'Z' || s[end] >= b'a' && s[end] <= b'z') {
                        end -= 1;
                    }
                    s.swap(start, end);
                    end -= 1;
                }
                start += 1;
            }
        }
        s
    }
}
