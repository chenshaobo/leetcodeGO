package threeSumCloest

import (
	"fmt"
	"math"
)


//Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
//Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//For example, given array S = {-1 2 1 -4}, and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

//1.遍历算出所有三个数相加和(和等于target则返回),否则然后找出最接近的数

func ThreeSumClosest(nums []int, target int) int {
	length := len(nums)
	quickSort(nums,0,length - 1)
	result := math.MaxInt32
	for i:=0 ; i< length;i++{
		h ,t := i+1,length-1
		for h < t{
			sum := nums[i] + nums[h] + nums[t]
			fmt.Printf("sum:%v,i:%v,h:%v,t:%v,result:%v,%v,%v\n",sum,i,h,t,result,abs(sum - target),abs(result - target))
			if sum == target{
				return target
			}
			if abs(sum - target)  < abs(result - target){
				result = sum
			}

			if sum > target {
				t --
			}else if sum < target{
				h ++
			}
		}
	}
	return result
}


func abs(n int)int{
	fmt.Printf("n:%v,:%v,:%v\n",n,n < 0,-1 * n)
	if n < 0 {
		return -1 * n
	}
	return n
}
func quickSort(s []int, start, end int) {
	if end - start < 1 {
		return
	}
	midValue := s[start]

	head, tail := start + 1, end
	//遍历
	for head <= tail {
		if s[head] >= midValue {
			if s[head] > s[tail] {
				s[head], s[tail] = s[tail], s[head]
			}
			tail --
		} else {
			head ++
		}
	}
	s[head - 1], s[start] = s[start], s[head - 1]
	quickSort(s, start, head - 1)

	quickSort(s, head, end)
}
