func moveZeroes(nums []int)  {
    i, j, n := 0, 0, len(nums)
    for j < n {
        nums[i] = nums[j]
        j++
        if nums[i] != 0 {
            i++
        }
    }
    for i < n {
        nums[i] = 0
        i++
    }
}