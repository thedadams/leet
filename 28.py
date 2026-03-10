class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        j = 0
        i = 0
        if needle == "":
            return 0
        while i < len(haystack):
            if haystack[i] == needle[j]:
                j += 1
                if j == len(needle):
                    return i - len(needle) + 1
            else:
                i -= j
                j = 0
            i += 1
        return -1