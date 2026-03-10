class Solution:
    def reverse(self, x: int) -> int:
        overflow = 2**31 - 1
        negative = x < 0
        if negative:
            x = -x
        y = 0
        while x > 0:
            y = (x % 10) + (y * 10)
            x //= 10
            if y > overflow:
                return 0
        return -y if negative else y
            
        