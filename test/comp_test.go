package comp_test

import (
	"fmt"
	"testing"

	"github.com/buguang01/comp"
)

func TestComp(t *testing.T) {
	s := "(a+b*c-4)/self.ID"
	cp := comp.NewCompNode(s)
	o := &SelfObj{}
	et := &comp.CompEvent{
		ObjList: []comp.Iobject{comp.NewObjectAttr("self", o)},
		Param:   []int64{10, 2, 7},
	}
	fmt.Println(cp.CompVal(et))
	fmt.Println(cp)
}

type SelfObj struct {
}

func (obj *SelfObj) GetAttr(attr string) int64 {
	return 10
}

func BenchmarkComp(b *testing.B) {
	s := "(a+b*c-4)/self.ID"
	cp := comp.NewCompNode(s)
	o := &SelfObj{}

	for i := 0; i < b.N; i++ {
		et := &comp.CompEvent{
			ObjList: []comp.Iobject{comp.NewObjectAttr("self", o)},
			Param:   []int64{10, 2, 7},
		}
		_ = cp.CompVal(et)
	}
}
