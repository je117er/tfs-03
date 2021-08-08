package utils

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) TopValue() (value interface{}) {
	return s.top.value
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Clear() (value interface{}) {
	if s.size > 0 {
		value, s.top = nil, nil
		s.size = 0
	}
	return nil
}

/*
func main() {
	stack := new(Stack)
	stack.Push("Things")
	stack.Push("and")
	stack.Push("Stuff")
	fmt.Printf("%s\n", stack.TopValue().(string))
	for stack.Len() > 0 {
		fmt.Printf("%s ", stack.Pop().(string))
	}
	stack.Clear()
	fmt.Println(stack.Len())
}

*/