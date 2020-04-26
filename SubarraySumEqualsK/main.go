func subarraySum(nums []int, k int) int {
    n := len(nums)
    var count int
    for i := 0; i < n; i++ {
        var sum int
        for j := i; j < n; j++ {
            sum += nums[j]
            if sum == k {
                count++
            }
        }
    }
    return count
}