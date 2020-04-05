def move_zeroes(nums)
    i = 0
    j = 0
    n = nums.length
    while j < n
        nums[i] = nums[j]
        j += 1
        i += 1 if nums[i] != 0
    end
    while i < n
        nums[i] = 0
        i += 1
    end
end