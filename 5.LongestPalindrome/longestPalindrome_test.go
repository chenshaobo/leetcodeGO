package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var s,s1 string
	s = "abcddcba"
	s1 = LongestPalindrome(s)
	if s != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}

	s = "caba"
	s1 = LongestPalindrome(s)
	if "aba" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}


	s = "bb"
	s1 = LongestPalindrome(s)
	if "bb" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}
}
