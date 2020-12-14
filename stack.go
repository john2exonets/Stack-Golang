package stackany

//  Stack.go -- a LIFO Stack
//
//  s.IsEmpty() () : true ? false
//  s.Push(v)	(a -- a b)
//  s.Pop() 	(a b -- b) : a
//  s.Drop()	(a b -- b)
//  s.Peek()	(a b -- a b) : a
//  s.Dup()		(a b -- a a b)
//  s.Swap()	(a b -- b a)
//  s.Rot()		(a b c -- b c a)
//  s.Over()	(a b -- a b a)
//  s.Pick(n)	(a b ... n -- n a b)
//  s.Nip()		(a b c -- a c)  {Swap; Drop;}
//  s.Tuck()	(a b -- b a b)	{Swap; Over;}
//  s.Move(&s2)	( a b : e f -- b : a e f)  {Pop; Push}
//  s.Dump()	(a) : Prints stack on stdout
//  s.Depth()	(a b c -- a b c) : 3
//  s.Clear()	(a b c --  )
//
//  John D. Allen
//  December, 2020
//

import "fmt"

// Stack holds the stack
type Stack []interface{}

// IsEmpty checks to see if stack has no elements
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push puts the value at the top of the stack
// b -- a b
func (s *Stack) Push(value interface{}) {
	t := *s
	x := Stack{value}
	*s = append(x, t...)
}

// Pop the top value off the stack
// a b c -- b c
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	n := (*s)[0]
	*s = (*s)[1:]
	return n
}

// Drop the top value off the stack
// a --
func (s *Stack) Drop() {
	if s.IsEmpty() {
		return
	}
	*s = (*s)[1:]
}

// Peek at the top value of the stack
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return ""
	}
	return (*s)[0]
}

// Dup the top value and push it on the top of the stack
// a -- a a
func (s *Stack) Dup() {
	if s.IsEmpty() {
		return
	}
	n := (*s)[0]
	s.Push(n)
}

// Swap reverses the first two elements on the stack
//  a b -- b a
func (s *Stack) Swap() {
	if len(*s) < 2 { // Nothing to swap!
		return
	}
	n1 := (*s)[0]
	n2 := (*s)[1]
	t := (*s)[2:]
	*s = Stack{n2}
	*s = append(*s, n1)
	*s = append(*s, t...)
}

// Rot takes the top item and puts it in the third place
// a b c -- b c a
func (s *Stack) Rot() {
	if len(*s) < 3 { // Nothing to rotate!
		return
	}
	n1 := (*s)[0]
	n2 := (*s)[1]
	n3 := (*s)[2]
	t := (*s)[3:]
	*s = Stack{n2}
	*s = append(*s, n3)
	*s = append(*s, n1)
	*s = append(*s, t...)
}

// Over takes the top item on the stack and copies it to the third place on the stack
// a b -- a b a
func (s *Stack) Over() {
	if len(*s) < 2 { // Nothing to Over!
		return
	}
	n1 := (*s)[0]
	n2 := (*s)[1]
	t := (*s)[2:]
	*s = Stack{n1}
	*s = append(*s, n2)
	*s = append(*s, n1)
	*s = append(*s, t...)
}

// Pick the n'th element and copy to the top of the stack
// a b ... n -- n a b
func (s *Stack) Pick(n int) {
	if len(*s) < n { // Stack is shorter than n!
		return
	}
	v := (*s)[n]
	s.Push(v)
}

// Nip does a Swap then Drop
// a b c -- a c
func (s *Stack) Nip() {
	if len(*s) < 2 { // Nothing to Nip!
		return
	}
	s.Swap()
	s.Drop()
}

// Tuck does a Swap then Over
// a b -- b a b
func (s *Stack) Tuck() {
	if len(*s) < 2 { // Nothing to Tuck!
		return
	}
	s.Swap()
	s.Over()
}

// Move takes the top item and moves it to the other stack
// a b : a b -- b : a a b
// Called like:  s.Move(&s2)
func (s *Stack) Move(s2 *Stack) {
	s2.Push(s.Pop())
}

// Dump out the whole stack
func (s *Stack) Dump() {
	for _, value := range *s {
		fmt.Println(value)
	}
}

// Depth of the stack
func (s *Stack) Depth() int {
	return len(*s)
}

// Clear the stack and reset to Empty
func (s *Stack) Clear() {
	*s = Stack{}
}
