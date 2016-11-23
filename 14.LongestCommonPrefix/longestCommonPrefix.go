package longestCommonPrefix


func LongestCommonPrefix(strs []string) string {
	strNumbers :=  len(strs)
	if strNumbers == 0 {
		return ""
	}
	if strNumbers == 1{
		return strs[0]
	}
	var s []byte
	for i:= 0;i< strNumbers; i ++ {
		if i==0{
			s = getCommonPrefix([]byte(strs[i]),[]byte(strs[i+1]))
			i++
		}else{
			s = getCommonPrefix(s,[]byte(strs[i]))
		}
	}
	return string(s)
}

func getCommonPrefix(a,b []byte) []byte{

	l := min(len(a),len(b))
	i:=0
	for ;i<l;i++{
		if a[i] != b[i]{
			break
		}
	}
	return a[:i]
}

func min(a,b int) int{
	if a > b {
		return b
	}
	return a
}