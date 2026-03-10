class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        found_nums = dict()
        for i, num in enumerate(nums):
            if num in found_nums:
                return [found_nums[num], i]
            found_nums[target - num] = i