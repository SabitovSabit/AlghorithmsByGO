package main

import (
	"fmt"
	
)

func main(){

  var searchNumber int
  fmt.Println("Enter Number!")
  fmt.Scan(&searchNumber)

  var res=fibonnaciRecursion(searchNumber,0,1)
 
  fmt.Printf("result:%v",res)
}