func rangeBitwiseAnd(m int, n int) int {
    for m < n {
        n -= n & -n
    }
    return n
}