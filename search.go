package main

var numbers = []int{1,3, 5, 5,6,1, 7, 8,9,1,5}

func linearSrch(num int)int {
	   for i:=0;i<len(numbers);i++{
		 if numbers[i]==num{
	         return i
		 }
	   }
	return -1
}

func findMaxNum()int {
	var max=numbers[0]

	for _,num:=range numbers{
        if num>=max{
			max=num
		}
	}
	return max
}

func findMostRepeated()int {
	var dic=make(map[int]int)
	var mostRepeatNum int
	var max=1

	for j:=0;j<len(numbers);j++{
        var count=0
		for i:=0;i<len(numbers);i++{
			if numbers[j]==numbers[i]{
				count+=1
			}
		}
		dic[numbers[j]]=count
	}

	for key,value:= range dic{
       if value>=max{
		 max =value
		 mostRepeatNum=key
	   }
	}

	return mostRepeatNum
}