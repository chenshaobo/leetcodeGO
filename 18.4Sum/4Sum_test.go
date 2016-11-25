package fourSum

import "testing"

func TestFourSum(t *testing.T) {
	var s []int
	var r [][]int
	s = []int{-3,-2,-1,0,0,1,2,3}
	r = FourSum(s,0)

	t.Logf("r:%v\n",r)
}
