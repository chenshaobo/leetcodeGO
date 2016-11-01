package reverseInteger

import (
	"strconv"
	"fmt"
)

func Reverse(x int) int {
	c := make([]byte,0)
	if x < 0 {
		c = append(c,'-')
		x = -1 * x
	}
	for {
		if x < 10{
			char := byte(x + '0')
			fmt.Printf("final char:%v\n",char)
			c = append(c, char)
			break
		}

		num := x % 10
		char := byte(num +'0')
		fmt.Printf(" char:%v\n",char)
		c = append(c,char)
		x = x/10
	}
	fmt.Printf("c:%v\n",c)
	i,e := strconv.ParseInt(string(c),10,32)
	if e != nil {
		return 0
	}
	fmt.Printf("e:%v\n",e)
	return int(i)
}