package main

var unNumbers = []int{3,1, 5, 45,6,67, 7, 8,9,0,13}

func selectionSort() []int{
   var temp int
 for j:=0;j<len(unNumbers);j++{

	 for i:=0;i<len(unNumbers)-1;i++{
		 if unNumbers[j]>unNumbers[i]{
			 temp=unNumbers[j]
			 unNumbers[j]=unNumbers[i]
			 unNumbers[i]=temp
			}
		}
	}
  return unNumbers
}

func bubbleSort() []int{

	var temp int
	for j:=0;j<len(unNumbers)-1;j++{

		for i:=0;i<len(unNumbers)-1-j;i++{
			if unNumbers[i]>unNumbers[i+1]{
				temp=unNumbers[i]			
				unNumbers[i]=unNumbers[i+1]
				unNumbers[i+1]=temp;
			}
	   }
	}
	 return unNumbers
}