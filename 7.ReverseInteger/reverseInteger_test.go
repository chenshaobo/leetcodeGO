package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
	i := 123
	ri := Reverse(i)
	if ri != 321{
		t.Fatalf("error:%v,%v",i,ri)
	}
}
