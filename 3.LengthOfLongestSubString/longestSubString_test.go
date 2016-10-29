package LengthOfLongestSubString

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	//subString,l := LengthOfLongestSubstring("pwwkew")
	//if subString != "wke" || l != 3 {
	//	t.Fatalf("sub:%v,len:%v\n", subString, l)
	//}
	//subString,l = LengthOfLongestSubstring("au")
	//if subString != "au" || l != 2 {
	//	t.Fatalf("sub:%v,len:%v\n", subString, l)
	//}




	subString,l := LengthOfLongestSubstring("dvdf")
	if subString != "vdf" || l != 3 {
		t.Fatalf("sub:%v,len:%v\n", subString, l)
	}
}
