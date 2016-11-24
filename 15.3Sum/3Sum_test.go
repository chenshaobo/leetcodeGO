package threeSum

import "testing"

func TestThreeSum(t *testing.T) {
	var s []int
	var r [][]int
	//s =  []int{0,0,0}
	//
	//r = ThreeSum(s)
	//
	//t.Logf("r :%v",r)
	//
	//s =  []int{1,2,-2,-1}
	//
	//r = ThreeSum(s)
	//
	//t.Logf("r :%v",r)

	s = []int{-1,0,1,2,-1,-4}

	r = ThreeSum(s)

	t.Logf("r :%v",r)
}
