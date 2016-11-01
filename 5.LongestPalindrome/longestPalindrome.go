package longestPalindrome

import "fmt"

/*
Given a string S, find the longest palindromic substring in S.
You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
 */
// abcddcbac
// caba
func cd LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	var max string = string(s[0])
	length := len(s)
	for i := 0; i < length - 1; i++ {
		left := i
		right := i

		fmt.Printf("===start with i=%v,max:%v===\n", i, max)
		for (left >= 0 && right < length) &&(s[left] == s[right] || s[i] == s[left] || s[right]== s[i]) {
			fmt.Printf("left:%v,right:%v------\n", left, right)
			if (s[left] == s[right]) {
				left--
				right ++
			} else if s[i] == s[left] {
				left --
			} else if s[i] == s[right] {
				right ++
			} else {
				break
			}
		}
		left ++
		right --
		if left < 0{
			left = 0
		}
		if right > length{
			right = length
		}
		if len(max) < right - left+1 {
			fmt.Printf("------left:%v,right:%v max:%v\n", left, right,s[left:right+1])
			max = s[left:right+1]
		}
	}
	return max

}
