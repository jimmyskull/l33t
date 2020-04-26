func rangeBitwiseAnd(m int, n int) int {
    var result int
    for m > 0 && n > 0 {
        mMSB := mostSignificantBit(m)
        nMSB := mostSignificantBit(n)
        if mMSB != nMSB {
            break
        }
        val := 1 << mMSB
        result += val
        m -= val
        n -= val
    }
    return result
}

func mostSignificantBit(m int) int {
    msb := -1
    for m != 0 {
        m >>= 1
        msb++
    }
    return msb
}