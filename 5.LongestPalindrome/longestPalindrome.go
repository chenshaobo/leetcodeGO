package longestPalindrome

//import "fmt"

/*
Given a string S, find the longest palindromic substring in S.
You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
 */
// abcddcbac
// caba
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) ==2{
		if s[0] == s[1]{
			return  s
		}
		return string(s[0])
	}
	var max string = string(s[0])
	length := len(s)
	for i := 0; i < length - 1; i++ {
		//fmt.Printf("===start with i=%v,max:%v===\n", i, max)
		s1 :=get(s,i,i,i,&max)
		if s1 !=""{
			max = s1
		}
		s2 := get(s,i,i,i+1,&max)
		if s2 !=""{
			max = s2
		}
	}
	return max
}


func get(s string,i ,left,right int,max * string) string{
	length := len(s)
	for (left >= 0 && right <= length -1) && (s[left] == s[right]) {
		//fmt.Printf("1-----left:%v,right:%v,i:%v\n", left, right,i)
		left --
		right ++
	}
	if len(*max) < right -left -1{
		//fmt.Printf("2-----left:%v,right:%v,i:%v\n", left, right,i)
		if right >= length{
			right = length
		}

		return s[left+1:right]

	}
	return ""
}