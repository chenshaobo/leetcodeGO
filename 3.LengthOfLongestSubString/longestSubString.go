package LengthOfLongestSubString

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
func LengthOfLongestSubstring(s string) (string,int ){
	if s == ""{
		return "",0
	}
	b := []byte(s)
	min := 0
	max := 1
	maxLen := 0
	j := 0
	i := 0
	m := map[byte]int{}
	for  ; i < len(b); i++{
		fmt.Printf("i:%v,j:%v\n",i,j)
		if found,ok := m[s[i]]; ok {
			fmt.Printf("found:%v,i:%v,j:%v\n",found,i,j)
			j = getMax(j,found +1)

		}
		m[s[i]] = i
		fmt.Printf("i:%v,j:%v\n",i,j)
		maxLen = getMax(maxLen,i-j+1)
	}
	min = j
	max = i
	fmt.Printf("min:%v,max:%v\n", min, max)
	var subString string
	if min == 0 && max == 0 {
		subString = string(b[0])
	}
	subString = string(b[min:max])
	return subString, maxLen
	//return maxLen
}


func getMax(i,j int) int{
	if i >= j {
		return i
	}
	return j
}