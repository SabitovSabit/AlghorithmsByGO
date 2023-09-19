package main

import (
	"fmt"
	
)

func main(){

  var searchNumber int
  fmt.Println("Enter number!")
  fmt.Scan(&searchNumber)

  var res=findMostRepeated()
 
  fmt.Printf("result:%v",res)
}