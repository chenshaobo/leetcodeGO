package medianOfTwoSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {


	nums1 := []int{1,2,3,4}
	nums2 := []int{5,6,7,8}
	r:=FindMedianSortedArrays(nums1,nums2)
	if 4.5 != r{
		t.Fatalf("Not Right :%v,%v,%v",nums1,nums2,r)
	}
}
