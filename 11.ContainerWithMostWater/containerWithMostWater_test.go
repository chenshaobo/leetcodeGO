package containerWithMostWater

import "testing"

func TestMaxArea(t *testing.T) {
	v := []int{2,3,4,5,18,17,6}
	r := MaxArea(v)
	if 17 != r {
		t.Fatalf("error:%v",r)
	}
}
