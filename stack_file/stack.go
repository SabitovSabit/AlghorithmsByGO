package stack_pck


type Stack struct{
	items[] int
}

func (s *Stack) Push(num int){
  s.items=append(s.items,num)
}

func (s *Stack) Pop() int{
  removeItem:= s.items[len(s.items)-1]
  s.items=s.items[:len(s.items)-1]
  return removeItem
}

func  (s *Stack) ReturnAll() []int {
	
	return s.items
}