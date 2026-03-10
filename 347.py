import heapq

class NumWithFreq:
    def __init__(self, num, freq):
        self.num = num
        self.freq = freq
    
    def add_one(self):
        self.freq += 1
    
    def __lt__(self, other):
        return self.freq < other.freq

    def __eq__(self, other):
        return self.num == other.num

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        d = dict()
        heap = []
        for num in nums:
            if num not in d:
                d[num] = NumWithFreq(num, 0)
            d[num].add_one()
        values = list(d.values())
        self.quickselect(values, 0, len(values) - 1, k)
        return [value.num for value in values[-k:]]
        

    def quickselect(self, arr, start, end, k):
        pivot = arr[end]
        index = start
        i = start
        while i <= end:
            if arr[i] < pivot:
                temp = arr[i]
                arr[i] = arr[index]
                arr[index] = temp
                index += 1
            i += 1
        temp = arr[index]
        arr[index] = pivot
        arr[end] = temp
        if index < len(arr) - k:
            self.quickselect(arr, index + 1, end, k)
        if index > len(arr) - k:
            self.quickselect(arr, start, index - 1, k)
            