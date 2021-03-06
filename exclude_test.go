package ndn

import (
	"reflect"
	"testing"

	"github.com/NDNLink/lpm"
)

func TestExclude(t *testing.T) {
	ex1 := Exclude{
		{Any: true},
		{Component: lpm.Component("AB")},
	}

	for _, test := range []struct {
		in   string
		want bool
	}{
		{"AB", true},
		{"AA", true},
		{"ABC", false},
	} {
		got := ex1.Match(lpm.Component(test.in))
		if got != test.want {
			t.Fatalf("Match(%v) == %v, got %v", test.in, test.want, got)
		}
	}

	b, err := ex1.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	var ex2 Exclude
	err = ex2.UnmarshalBinary(b)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ex1, ex2) {
		t.Fatalf("expect %+v, got %+v", ex1, ex2)
	}
}
