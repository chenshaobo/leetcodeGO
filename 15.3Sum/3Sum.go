package threeSum

import (
	"fmt"
)

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	var m = make(map[string]interface{})

	length := len(nums)
	quickSort(nums, 0, length - 1)
	fmt.Printf("sort list :%v\n", nums)

	if nums[0] > 0 {
		return result
	}
	negativeTail := 0
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		negativeTail ++
	}
	for i := 0; i < negativeTail; i++ {
		sum := nums[i]
		r := findAdded(m, sum, nums, i + 1, length - 1)
		result = append(result, r...)
	}
	return result
}

func findAdded(m map[string]interface{}, sum int, nums []int, start, end int) [][]int {
	postiveSum := - sum
	fmt.Printf("sum:%v,start:%v,end:%v\n", sum, start, end)
	result := make([][]int, 0)
	for i := start; i < end; i++ {
		want := postiveSum - nums[i]
		isFound := binarySearch(want, nums, i + 1, end)

		if isFound {
			fmt.Printf("s:%v,end:%v",start,end)
			rstr := fmt.Sprint(sum, nums[i], want)
			fmt.Printf("rstr:%v,map:%v\n", rstr, m)
			if _, ok := m[rstr]; !ok {
				r := []int{sum, nums[i], want}
				result = append(result, r)
				var tmp interface{}
				m[rstr] = tmp
			}

		}
	}
	return result
}

func binarySearch(want int, nums []int, start, end int) bool {
	fmt.Printf("binarySearch:want:%v,start:%v,end:%v\n", want, start, end)
	for start <= end {
		mid := int((start + end) / 2)
		fmt.Printf("start:%v,end:%v,mid:%v,want:%v,nums[mid]:%v,:%v\n", start, end, mid, want, nums[mid], start <= end)
		if start == mid || end == mid {
			if want == nums[start] {
				return true
			}
			if want == nums[end] {
				return true
			}
			break
		}
		if want == nums[mid] {
			return true
		}

		if want > nums[mid] {
			start = mid
		}

		if want < nums[mid] {
			end = mid
		}
	}

	return false
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
