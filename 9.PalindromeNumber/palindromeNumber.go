package palindromeNumber

import "fmt"

func IsPalindrome(x int) bool {
	r := 0
	t := x
	for {
		if t > 0 {
			b := t % 10
			r = r * 10 + b
			t = t / 10
			fmt.Printf("x:%v,r:%v\n",t,r)
		}else{
			break
		}
	}
	fmt.Printf("x:%v,r:%v\n",x,r)
	return x == r
}