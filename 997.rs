impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut trusts = vec![0; n as usize];
        for t in trust.into_iter() {
            trusts[(t[0]-1) as usize] -= 1;
            trusts[(t[1]-1) as usize] += 1;
        }

        for (i, t) in trusts.into_iter().enumerate() {
            if t == n - 1 {
                return (i + 1) as i32;
            }
        }
         -1
    }
}
