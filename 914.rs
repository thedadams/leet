impl Solution {
    pub fn has_groups_size_x(deck: Vec<i32>) -> bool {
        let mut counts: [i32; 10000] = [0; 10000];

        for card in deck.into_iter() {
            counts[card as usize] += 1;
        }

        let mut min_count: i32 = -1;
        for card in counts.into_iter() {
            min_count = Self::gcd(min_count, card);
        }

        min_count > 1
    }

    fn gcd(mut n: i32, mut m: i32) -> i32 {
        if m == 0 {
            return n;
        }
        if n == -1 {
            return m;
        }
        if n > m {
            std::mem::swap(&mut n, &mut m);
        }

        let mut r: i32 = m % n;
        while r > 0 {
            m = n;
            n = r;
            r = m % n;
        }

        n
    }
}
