package fourSum

import "fmt"

func FourSum(nums []int, target int) [][]int {
	fmt.Printf("nums:%v\n",nums)
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	var result [][]int
	quickSort(nums,0,length -1)
	fmt.Printf("sort nums:%v\n",nums)
	 m := make(map[string] interface{})
	for i:=0;i<length-1;i++{
		for j := i+1;j < length -1 ;j++{
			e1 := nums[i]
			e2 := nums[j]
			head,tail := j+1,length-1
			for head < tail {
				sum := e1 +e2 +nums[head] + nums[tail]

				if sum == target{
					str := fmt.Sprint(e1,e2,nums[head],nums[tail])
					_,ok := m[str]
					if !ok {
						result = append(result,[]int{e1,e2,nums[head],nums[tail]})
						var t interface{}
						m[str] = t
					}
					head ++
					tail --
				}else if sum > target{
					tail --
				}else if sum < target{
					head ++
				}
			}
		}


	}
	return result
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