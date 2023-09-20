package main

import (
	"fmt"
	
)

func main(){

  var searchNumber int
  fmt.Println("Enter Number!")
  fmt.Scan(&searchNumber)

  var res=bubbleSort()
 
  fmt.Printf("result:%v",res)
}