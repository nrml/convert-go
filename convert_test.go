package convert

import (
	"testing"
)

type T1 struct {
	I int8
	S string
	B bool
}
type T2 struct {
	I int64
	S string
	B bool
}

func Test_Convert(t *testing.T) {
	t1 := T1{5, "t1", true}
	t2 := new(T2)

	Convert(t1, t2)
	t.Logf("t2: %v\n", t2)
	if t2.I != 5 {
		t.Error("did not convert")
	}

}

func Test_ConvertMap(t *testing.T) {
	mp := make(map[interface{}]interface{})
	mp["I"] = 5
	mp["S"] = "m1"
	mp["B"] = true

	t2 := new(T2)
	ConvertMap(mp, t2)
	t.Logf("t2: %v\n", t2)
	if t2.I != 5 {
		t.Error("did not convert")
	}
}
