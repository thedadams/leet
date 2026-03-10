/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    start := 1
    var mid, ans int
    for start < n {
        mid = (start + n) / 2
        ans = guess(mid)
        if ans == 0 {
            return mid
        }
        if ans < 0 {
            n = mid - 1
        } else {
            start = mid + 1
        }
    }

    return start
}