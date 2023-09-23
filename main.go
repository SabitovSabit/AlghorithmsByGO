package main

import (
	"alghorithms/queue_file"
	"fmt"
)

func main(){

  var searchNumber int
  fmt.Println("Enter Number!")
  fmt.Scan(&searchNumber)

  myQueue:=queue_pck.Queue{}
  myQueue.Enqueue(searchNumber)
  myQueue.Enqueue(searchNumber+1)
  myQueue.Enqueue(searchNumber-1)

  myQueue.Dequeue()
  var res=myQueue.ReturnAll()
  
  fmt.Printf("result:%v",res)
}