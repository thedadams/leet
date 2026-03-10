class Solution:
    def reverseBits(self, n: int) -> int:
        y = 0
        for _ in range(32):
            y = (n % 2) + y << 1
            n >>= 1
        return y >> 1