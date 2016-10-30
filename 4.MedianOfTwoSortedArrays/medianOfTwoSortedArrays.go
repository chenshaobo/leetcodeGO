package medianOfTwoSortedArrays

import "fmt"


/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	maxLen := len1+len2
	s := make([]int,maxLen)
	j := 0
	i := 0
	fmt.Printf("maxLen:%v\n",maxLen)
	index :=0
	for {

		if index > maxLen{
			fmt.Printf("maxLen-1/2:%v",maxLen-1/2)
			if maxLen % 2 != 0 {
				return float64(s[(maxLen-1)/2])
			}
			return float64(s[(maxLen-1)/2] + s[(maxLen-1)/2+1]) / 2.0
		}else if i == len1{
			for ;j<len2;j++{
				fmt.Printf("j:%v,len:%v\n",j,len2)
				s[index] = nums2[j]
				index++
			}
		}else if j == len2{
			for ;i <len1;i++{
				s[index] = nums1[i]
				index ++
			}
		}else{
			if nums1[i] > nums2[j]{
				s[index] = nums2[j]
				j++
			}else{
				s[index] = nums1[i]
				i++
			}

		}
		fmt.Printf("index:%v,i:%v,j:%v.s:%v\n",index,i,j,s)
		index ++
		//fmt.Printf("index:%v,i:%v,j:%v.s:%v\n",index,i,j,s)

	}
}


func bigger(a,b int) int{
	if a > b {
		return a
	}
	return b
}