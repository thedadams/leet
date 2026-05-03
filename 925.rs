impl Solution {
    pub fn is_long_pressed_name(name: String, typed: String) -> bool {
        let nameBytes = name.as_bytes();

        let mut i = 0;
        let mut c = nameBytes[0];
        for &b in typed.as_bytes() {
            if i < nameBytes.len() && b == nameBytes[i] {
                i += 1;
            } else if b != nameBytes[i.saturating_sub(1)] {
                return false;
            }
        }

        i == nameBytes.len()
    }
}
