package threeSumCloest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	var s []int
	var r int

	s = []int{-1,2,1,-4}

	r = ThreeSumClosest(s,1)

	if r != 2 {
		t.Fatalf("Cloest not right:%v",r)
	}

	s = []int{-3,-2,-5,3,-4}

	r = ThreeSumClosest(s,-1)

	if r != -2 {
		t.Fatalf("Cloest not right:%v",r)
	}
}
