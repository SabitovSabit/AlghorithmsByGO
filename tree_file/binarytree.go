package tree_pck

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(num int) {

	if num > n.Data {
		if n.Right == nil {
			n.Right = &Node{Data: num}

		} else {
			n.Right.Insert(num)
		}

	} else if num < n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: num}
		} else {
			n.Left.Insert(num)
		}
	}
}

func (n *Node) Search(num int) int{
  
  if n==nil{
    return -1
  }

  if num > n.Data {
    return n.Right.Search(num)
  }else if num < n.Data{
    return n.Left.Search(num)
  }

  return n.Data
}
