/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    start, end, next := 1, n, 0

    for start < end {
        next = (end + start) / 2
        if isBadVersion(next) {
            end = next
        } else {
            start = next + 1
        }
    }

    return start
}