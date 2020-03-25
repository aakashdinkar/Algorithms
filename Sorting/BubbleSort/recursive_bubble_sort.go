package main

import "fmt"

func bubble_Sort(arr[] int, n int){
	var i int
	if n == 1{
		return
	}
	for i=0;i<n-1;i++{
			if arr[i] > arr[i+1]{
				arr[i],arr[i+1] = arr[i+1],arr[i]
				bubble_Sort(arr,n-1)
			}
	}
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
	bubble_Sort(arr,n)
	fmt.Println("Sorted array is :",arr)
}
