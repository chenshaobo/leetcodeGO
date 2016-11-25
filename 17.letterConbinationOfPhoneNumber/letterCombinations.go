package letterCombinations


func LetterCombinations(digits string) []string{
	if digits == ""{
		return []string{}
	}
	 m := map[byte][]string{
		 '2':[]string{"a","b","c"},
		 '3':[]string{"d","e","f"},
		 '4':[]string{"g","h","i"},
		 '5':[]string{"j","k","l"},
		 '6':[]string{"m","n","o"},
		 '7':[]string{"q","p","r","s"},
		 '8':[]string{"t","u","v"},
		 '9':[]string{"w","x","y","z"},
	 }
	result,ok := m[digits[0]];
	//fmt.Printf(":%v,%v\nm:%v\n,digits[0]:%v\n",result,ok,m,digits[0])
	if !ok{
		return []string{}
	}
	for i:=1;i< len(digits) ;i++{
		r,ok :=m[digits[i]]
		if !ok{
			return []string{}
		}
		//fmt.Printf("i:%v\n",m[digits[i]])
		newResult := make([]string,0)
		for _,v := range result{
			for _,
			v1 := range r{
				//fmt.Printf("newresult:%v,v:%v,v1:%v\n",newResult,v,v1)
				newResult = append(newResult,string(v)+string(v1))
			}
		}
		if len(newResult) > 0{
			result = newResult
		}
		
	}
	return result
}
