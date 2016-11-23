package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	s := []string{
		"leetcode",
		"leetcade",
		"leetc",
		"leet",
		"let",
		"lecc",
	}

	commonPrefix := LongestCommonPrefix(s)
	if "le" != commonPrefix{
		t.Fatalf("Not Right:%v,slice:%v\n",commonPrefix,s)
	}
}
