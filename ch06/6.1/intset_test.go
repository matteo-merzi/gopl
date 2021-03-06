package intset

import "testing"

func TestLen(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Len() != 3 {
		t.Errorf("s.Len() should be 3")
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Len() != 3 {
		t.Errorf("s.Len() should be 3")
	}
	s.Remove(1)
	if s.Len() != 2 {
		t.Errorf("s.Len() should be 2")
	}
	if s.String() != "{2 3}" {
		t.Errorf("s.String() should be {2 3}")
	}
}

func TestClear(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Len() != 3 {
		t.Errorf("s.Len() should be 3")
	}
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("s.Len() should be 0")
	}
	if s.String() != "{}" {
		t.Errorf("s.String() should be {}")
	}
}

func TestCopy(t *testing.T) {
	var a IntSet
	a.Add(1)
	a.Add(2)
	a.Add(3)
	if a.Len() != 3 {
		t.Errorf("a.Len() should be 3")
	}
	b := a.Copy()
	if b.Len() != 3 {
		t.Errorf("b.Len() should be 3")
	}
	if b.String() != "{1 2 3}" {
		t.Errorf("b.String() should be {1 2 3}")
	}
}
