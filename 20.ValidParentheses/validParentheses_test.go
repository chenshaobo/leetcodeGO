package validParentheses

import "testing"

func TestIsValid(t *testing.T) {

	var s string
	s = "{}()[]"
	if IsValid(s) != true{
		t.Fatalf("Valid erro:%v\n",s)
	}

	s = "({[]})"
	if IsValid(s) != true{
		t.Fatalf("Valid erro:%v\n",s)
	}

	s = "[({(())}[()])]"
	if IsValid(s) != true{
		t.Fatalf("Valid erro:%v\n",s)
	}
}
