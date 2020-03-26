package main

import "fmt"

func selection_Sort(arr[] int, n int){
	var i,j,min_index int
	for i=0;i<n-1;i++{
		min_index = i
		for j=i+1;j<n;j++{
			if arr[j] < arr[min_index]{
				min_index=j
			}
		}
		arr[min_index],arr[i] = arr[i],arr[min_index]
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
	selection_Sort(arr,n)
	fmt.Println("Sorted array is :",arr)
}
