package main

import "fmt"

func main() {
	var r int
	r = longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	fmt.Println(r, 1)
	r = longestCommonSubsequence("abcde", "ace")
	fmt.Println(r, 3)
	r = longestCommonSubsequence("abc", "abc")
	fmt.Println(r, 3)
	r = longestCommonSubsequence("abc", "def")
	fmt.Println(r, 0)

}

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	if n == 0 || m == 0 {
		return 0
	}
	mat := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				mat[i*m+j] = getValue(mat, n, m, i-1, j-1) + 1
			} else {
				up := getValue(mat, n, m, i-1, j)
				left := getValue(mat, n, m, i, j-1)
				max := maxValue(up, left)
				mat[i*m+j] = max
			}
		}
	}
	return mat[n*m-1]
}

func getValue(mat []int, n, m, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return mat[i*m+j]
}

func maxValue(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
