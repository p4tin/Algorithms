package Stack

import "testing"

type tstruct struct {
	x int
	s string
}

func TestStack(t *testing.T) {
	s := CreateStack()
	if !s.IsEmpty() {
		t.Errorf("1 - Stack should be empty but isn't")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("2 - Stack is empty but should have an entry")
	}

	s.Push(2)

	i := s.Pop()
	if i == nil {
		t.Errorf("3 - Should get something of the stack but got <nil>")
	}
	if i != 2 {
		t.Errorf("4 - Expect 2 off the stack but got %d", i)
	}

	i = s.Pop()
	if i != 1 {
		t.Errorf("5 - Expect 1 off the stack but got %d", i)
	}

	if !s.IsEmpty() {
		t.Errorf("6 - Stack should be empty but isn't")
	}

	s.Push(true)
	i = s.Pop()
	if i != true {
		t.Errorf("7 - Expected true off the stack but got %b", i)
	}

	st := tstruct{
		s: "2",
		x: 1,
	}
	s.Push(st)
	st_r := s.Pop()
	if st_r.(tstruct).x != 1 {
		t.Errorf("8 - Expect 1 off the stack but got %d", st_r.(tstruct).x)
	}
}
