class Solution:
    def maxPower(self, s: str) -> int:
        this_char = ''
        this_streak = 0
        max_streak = 0
        for c in s:
            if c == this_char:
                this_streak += 1
            else:
                this_char = c
                max_streak = max(max_streak, this_streak)
                this_streak = 1
        return max(max_streak, this_streak)