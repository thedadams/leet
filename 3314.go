func minBitwiseArray(nums []int) []int {
    for i, num := range nums {
        if num == 2 {
            nums[i] = -1
            continue
        }

        pow2 := 1
        for pow2 < num {
            if num & (pow2 << 1) == 0 {
                nums[i] ^= pow2
                break
            }

            pow2 <<= 1
        }
        if pow2 > num {
            nums[i] >>= 1
        }
    }

    return nums
}