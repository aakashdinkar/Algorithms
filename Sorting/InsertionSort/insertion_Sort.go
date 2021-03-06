package main

import "fmt"

func insertion_Sort(arr[] int, n int){
	var i,j,key int
	for i = 1; i < n; i++{
		key = arr[i]
		j = i - 1
		for(j >= 0 && arr[j] > key){
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
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
	insertion_Sort(arr,n)
	fmt.Println("Sorted array is :",arr)
}
