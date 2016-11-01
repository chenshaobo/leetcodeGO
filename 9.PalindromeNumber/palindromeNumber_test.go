package palindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	a := 1234321
	r := IsPalindrome(a)
	if true != r{
		t.Fatalf("Error:%v,%v",a,r)
	}
}