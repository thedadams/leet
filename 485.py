class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        max_streak = 0
        this_streak = 0
        for num in nums:
            if num == 1:
                this_streak += 1
            else:
                max_streak = max(max_streak, this_streak)
                this_streak = 0
        return max(max_streak, this_streak)