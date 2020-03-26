package main

import "fmt"

func minindex(arr[] int,i,j int)int{
	if i == j{
		return i
	}
	k := minindex(arr,i+1,j)
	if arr[i] < arr[k]{
		return i
	}else{ return k
  }
}

func selection_Sort(arr[] int, n int, index int){
	if index == n{
		return
	}

	k := minindex(arr,index,n-1)
	if k != index{
		arr[k],arr[index] = arr[index],arr[k]
	}
	selection_Sort(arr,n,index+1)
}

func main(){
	var n int
	fmt.Print("Enter number of elements:")
	fmt.Scan(&n)
	fmt.Println("Enter the elements of array")
	arr := make([]int, n)
	i:=0
	for(i<n){
		fmt.Print("Enter element ",i+1," :")
		fmt.Scan(&arr[i])
		i = i + 1
	}
	selection_Sort(arr,n,0)
	fmt.Println("Sorted array is :",arr)
}
