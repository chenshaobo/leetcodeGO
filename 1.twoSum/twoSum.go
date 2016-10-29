package twoSum

import (
	"github.com/kataras/go-errors"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

//TwoSum must faster than TwoSum when elements's length is big enough
func TwoSum(elements []int,target int)[]int{
	m := make(map[int] int,len(elements))
	for i,v := range elements{
		if index2,ok := InSlice(elements,i,target - v);ok == nil{
			return []int{i,index2}
		}
		m[v] = i
	}
	for i,v := range elements{
		if index2,ok := m[target -v ];ok{
			return []int{i,index2}
		}
	}
	return []int{}
}



func InSlice(s []int,expect int,target int) (int,error){
	for i,v := range s{
		if v == target && i != expect {
			return i,nil
		}
	}
	return -1,errors.New("Not int slice")
}


func TwoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[n]; ok {
			return []int{j, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}
