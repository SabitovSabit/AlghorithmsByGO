package main

import (
	tree_pck "alghorithms/tree_file"
	"fmt"
)

func main() {

	/*var searchNumber int
	  fmt.Println("Enter Number!")
	  fmt.Scan(&searchNumber)

	  myQueue:=queue_pck.Queue{}
	  myQueue.Enqueue(searchNumber)
	  myQueue.Enqueue(searchNumber+1)
	  myQueue.Enqueue(searchNumber-1)

	  myQueue.Dequeue()
	  var res=myQueue.ReturnAll()

	  fmt.Printf("result:%v",res)*/

	 bts := &tree_pck.Node{Data:4}
 
   bts.Insert(3)
   bts.Insert(11)
   bts.Insert(33)
   bts.Insert(44)

  var src=bts.Search(33)

	fmt.Println(src)
}
