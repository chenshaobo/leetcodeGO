package LengthOfLongestSubString

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */
func LengthOfLongestSubstring(s string) (string, int) {
	b := []byte(s)
	min := 0
	max := 1
	maxLen := 1
	j := 0
	i := 0
	m := map[byte]int{}
	for ; i < len(b); i++ {
		m[s[i]] = i
		for ; j < len(b); j++ {
			v, ok := m[s[j]]
			fmt.Printf("i:%v,j:%v,%v:%v\n", i, j, s[j], v)
			add := 0
			if i == 0 {
				add = 1
			} else {
				add = 0
			}
			fmt.Printf("========i:%v,j:%v,%v:%v\n", i, j, s[j], m[s[j]])
			if j > v && ok && v !=0 {
				fmt.Printf("-----------j:%v,v:%v\n", j, v)
				if (j - i + add) >= maxLen {
					min = i
					max = j + add
					maxLen = max - min
					fmt.Printf("********i:%v,j:%v,s[j]:%v,%v,maxlen:%v\n", i, j, s[j], m[s[j]], maxLen)
				}
				break
			}
			if (j - i + add) >= maxLen {
				min = i
				max = j + add
				maxLen = max - min
				fmt.Printf("********i:%v,j:%v,s[j]:%v,%v,maxlen:%v\n", i, j, s[j], m[s[j]], maxLen)
			}
			m[s[j]] = j
		}
	}

	fmt.Printf("min:%v,max:%v", min, max)
	var subString string
	if min == 0 && max == 0 {
		subString = string(b[0])
	}
	subString = string(b[min:max])
	return subString, maxLen
}