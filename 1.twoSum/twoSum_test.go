package twoSum


import(
	"testing"
)


func TestTwoSum(t *testing.T) {
	s := []int{1,2,3,4,4,4,4,8,9,10}
	r := TwoSum(s,9)
	if r[0]!=0 || r[1] != 7{
		t.Fatalf("When input:%v,result is :%v not right ",s,r)
	}


	s = []int{3,2,4}
	r = TwoSum(s,6)
	if r[0]!=1 || r[1] != 2{
		t.Fatalf("When input:%v,result is :%v not right ",s,r)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	s := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
		1,1,1,1,1,1,1,1,1,1,11,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,12,3,4,5,6,7,8,9,10}
	for i:=0;i < b.N;i++{
		TwoSum(s,9)
	}
}

func TestTwoSum1(t *testing.T) {
	s := []int{1,2,3,4,4,4,4,8,9,10}
	r := TwoSum1(s,9)
	if r[0]!=0 || r[1] != 7{
		t.Fatalf("When input:%v,result is :%v not right ",s,r)
	}


	s = []int{3,2,4}
	r = TwoSum1(s,6)
	if r[0]!=1 || r[1] != 2{
		t.Fatalf("When input:%v,result is :%v not right ",s,r)
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	s := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
		1,1,1,1,1,1,1,1,1,1,11,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,12,3,4,5,6,7,8,9,10}
	for i:=0;i < b.N;i++{
		TwoSum1(s,9)
	}
}
