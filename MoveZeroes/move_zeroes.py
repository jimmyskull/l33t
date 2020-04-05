class Solution(object):
    
    def moveZeroes(self, nums):
        i, j, n = 0, 0, len(nums)
        while j < n:
            nums[i] = nums[j]
            j += 1
            if nums[i] != 0:
                i += 1
        while i < n:
            nums[i] = 0
            i += 1