package containerWithMostWater

import (
	"fmt"
	"math"
)

func MaxArea(height []int) int {
	l := len(height)
	j := l -1
	max := 0
	for i :=0;i<j ;i++ {
		//j := l-1
		i1 := i
		for i1 < j {
			fmt.Printf("i:%v,j:%v,ii:%v,jj:%v,r:%v\n",i1,j,height[i1],height[j],math.Min(float64(height[i1]),float64(height[j])) * float64(j - i1))
			if height[i1] > height[j] {
				area := (j - i1) * height[j]
				max = setMax(area, max)
				j--
			} else {
				area := (j - i1) * height[i1]
				max = setMax(area, max)
				i1++
			}
		}
	}
	return max
}

func setMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}