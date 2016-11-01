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

	s = "adam"
	s1 = LongestPalindrome(s)
	if "ada" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}

	s = "aaabaaaa"
	s1 = LongestPalindrome(s)
	if "aaabaaa" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}

	s = "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"
	s1 = LongestPalindrome(s)
	if "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}
	s = "eabcb"
	s1 = LongestPalindrome(s)
	if "bcb" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}

	s = "abb"

	s1 = LongestPalindrome(s)
	if "bb" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}

	s = "bba"

	s1 = LongestPalindrome(s)
	if "bb" != s1{
		t.Fatalf("Not Right :%v,%v",s,s1)
	}
}

