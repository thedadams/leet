type RecentCounter struct {
    calls []int
    smallestIdx int
}


func Constructor() RecentCounter {
    return RecentCounter{calls: make([]int, 0)}
}


func (r *RecentCounter) Ping(t int) int {
    for len(r.calls) > 0 && r.smallestIdx < len(r.calls) {
        if t > r.calls[r.smallestIdx] + 3000 {
            r.smallestIdx++
        } else {
            break
        }
    }

    r.calls = append(r.calls, t)

    return len(r.calls) - r.smallestIdx
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
