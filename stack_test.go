package stackany

import (
	"testing"

	tt "github.com/john2exonets/golang-testing-util-functions"
)

//------------------[   Tests   ]------------------------

func TestPush(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	s.Push("B")
	s.Push("C")
	// ["C","B","A"]
	tt.Contains(t, "B", s)
	v := s.Pop()
	tt.Assert(t, v, "C")
	v = s.Pop()
	tt.Assert(t, v, "B")
	v = s.Pop()
	tt.Assert(t, v, "A")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(3)
	s.Push(2)
	s.Push(1)
	// [1,2,3]
	i := s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 2)
	i = s.Pop()
	tt.Assert(t, i, 3)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestDrop(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	s.Push("B")
	// ["B","A"]
	s.Drop()
	v := s.Pop()
	tt.Assert(t, v, "A")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(1)
	s.Push(2)
	// [2,1]
	s.Drop()
	i := s.Pop()
	tt.Assert(t, i, 1)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestDup(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	s.Dup()
	v := s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "A")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(1)
	s.Dup()
	i := s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 1)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestPeek(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	v := s.Peek()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "A")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(1)
	i := s.Peek()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 1)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestEmpty(t *testing.T) {
	var s Stack
	tt.Assert(t, s.IsEmpty(), true)
	s.Push("A")
	tt.Assert(t, s.IsEmpty(), false)
	v := s.Pop()
	tt.Assert(t, v, "A")
	s.Push(1)
	tt.Assert(t, s.IsEmpty(), false)
	v = s.Pop()
	tt.Assert(t, v, 1)
}

func TestLen(t *testing.T) {
	var s Stack

	s.Push("A")
	s.Push("B")
	s.Push("C")
	s.Push("D")
	tt.Assert(t, s.Depth(), 4)
}

func TestSwap(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	s.Push("B")
	// ["B","A"]
	s.Swap()
	// ["A","B"]
	v := s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "B")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(1)
	s.Push(2)
	// [2,1]
	s.Swap()
	// [1,2]
	i := s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 2)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestRot(t *testing.T) {
	var s Stack

	// string tests
	s.Push("C")
	s.Push("B")
	s.Push("A")
	// [ "A","B","C"]
	s.Rot()
	// ["B","C","A"]
	v := s.Pop()
	tt.Assert(t, v, "B")
	v = s.Pop()
	tt.Assert(t, v, "C")
	v = s.Pop()
	tt.Assert(t, v, "A")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(3)
	s.Push(2)
	s.Push(1)
	// [ 1,2,3 ]
	s.Rot()
	// [ 2,3,1 ]
	i := s.Pop()
	tt.Assert(t, i, 2)
	i = s.Pop()
	tt.Assert(t, i, 3)
	i = s.Pop()
	tt.Assert(t, i, 1)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestOver(t *testing.T) {
	var s Stack

	// string tests
	s.Push("A")
	s.Push("B")
	// ["B","A"]
	s.Over()
	// ["B","A","B"]
	v := s.Pop()
	tt.Assert(t, v, "B")
	v = s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "B")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(1)
	s.Push(2)
	// [ 2,1 ]
	s.Over()
	// [ 2,1,2 ]
	i := s.Pop()
	tt.Assert(t, i, 2)
	i = s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 2)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestNip(t *testing.T) {
	var s Stack

	// string tests
	s.Push("C")
	s.Push("B")
	s.Push("A")
	// ["A","B","C"]
	s.Nip()
	// ["A","C"]
	v := s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "C")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(3)
	s.Push(2)
	s.Push(1)
	// [ 1,2,3 ]
	s.Nip()
	// [ 1,3 ]
	i := s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 3)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestTuck(t *testing.T) {
	var s Stack

	// string tests
	s.Push("B")
	s.Push("A")
	// ["A","B"]
	s.Tuck()
	// ["B","A","B"]
	v := s.Pop()
	tt.Assert(t, v, "B")
	v = s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "B")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(2)
	s.Push(1)
	// [ 1,2 ]
	s.Tuck()
	// [ 2,1,2 ]
	i := s.Pop()
	tt.Assert(t, i, 2)
	i = s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 2)
	tt.Assert(t, s.IsEmpty(), true)
}

func TestClear(t *testing.T) {
	var s Stack

	s.Push("A")
	s.Push("B")
	// ["B","A"]
	tt.Assert(t, s.Depth(), 2)
	s.Clear()
	// []
	tt.Assert(t, s.IsEmpty(), true)
}

func TestMove(t *testing.T) {
	var s1 Stack
	var s2 Stack

	// string tests
	s1.Push("B")
	s1.Push("A")
	// s1: ["A","B"]
	s2.Push("F")
	s2.Push("E")
	// s2: ["E", "F"]
	s1.Move(&s2)
	// s1: ["B"], s2: ["A","E","F"]
	v := s1.Pop()
	tt.Assert(t, v, "B")
	v = s2.Pop()
	tt.Assert(t, v, "A")
	v = s2.Pop()
	tt.Assert(t, v, "E")

	// int tests
	s1.Push(2)
	s1.Push(1)
	// s1: [ 1,2 ]
	s2.Push(6)
	s2.Push(5)
	// s2: [ 5,6 ]
	s1.Move(&s2)
	// s1: [ 2 ], s2: [ 1,5,6 ]
	i := s1.Pop()
	tt.Assert(t, i, 2)
	i = s2.Pop()
	tt.Assert(t, i, 1)
	i = s2.Pop()
	tt.Assert(t, i, 5)
}

func TestPick(t *testing.T) {
	var s Stack

	// string tests
	s.Push("C")
	s.Push("B")
	s.Push("A")
	// ["A","B","C"]
	s.Pick(2)
	// ["C","A","B","C"]
	v := s.Pop()
	tt.Assert(t, v, "C")
	v = s.Pop()
	tt.Assert(t, v, "A")
	v = s.Pop()
	tt.Assert(t, v, "B")
	v = s.Pop()
	tt.Assert(t, v, "C")
	tt.Assert(t, s.IsEmpty(), true)

	// int tests
	s.Push(3)
	s.Push(2)
	s.Push(1)
	// [ 1,2,3 ]
	s.Pick(2)
	// [ 3,1,2,3 ]
	i := s.Pop()
	tt.Assert(t, i, 3)
	i = s.Pop()
	tt.Assert(t, i, 1)
	i = s.Pop()
	tt.Assert(t, i, 2)
	i = s.Pop()
	tt.Assert(t, i, 3)
	tt.Assert(t, s.IsEmpty(), true)
}
