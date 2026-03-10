class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        zero_count = 0
        i = 0
        for num in nums:
            if num != 0:
                nums[i] = num
                i += 1
            else:
                zero_count += 1
        for j in range(zero_count):
            nums[i + j] = 0
        