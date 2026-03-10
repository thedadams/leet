func majorityElement(nums []int) int {
    var result, majority int
    for _, i := range nums {
        if majority == 0 {
            result = i
        }

        if result == i {
            majority++
        } else {
            majority--
        }
    }

    return result
}