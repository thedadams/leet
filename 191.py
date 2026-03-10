class Solution:
    def hammingWeight(self, n: int) -> int:
        one_bits = 0
        while n > 0:
            one_bits += 1 if n % 2 else 0
            n >>= 1
        return one_bits