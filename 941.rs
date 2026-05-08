impl Solution {
    pub fn valid_mountain_array(arr: Vec<i32>) -> bool {
        let mut decreasing = false;
        for i in 1..arr.len() {
            if arr[i-1] < arr[i] {
                if decreasing {
                    return false;
                }
            } else if arr[i-1] > arr[i] {
                if i == 1 {
                    return false;
                }
                decreasing = true
            } else {
                return false;
            }
        }

        decreasing
    }
}
