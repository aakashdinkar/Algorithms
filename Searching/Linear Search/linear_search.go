package main

import "fmt"

func linear_search(arr[] int, n int, num int) int {
	i := 0
	for(i<n){
		if arr[i] == num{
			return i
		}
		i = i + 1
	}
	return -1
}

func main(){
	var n,num int
	fmt.Print("Enter number of elements:")
	fmt.Scan(&n)
	fmt.Println("Enter the elements")
	arr := make([]int, n)
	i:=0
	for(i<n){
		fmt.Print("Enter element ",i+1," :")
		fmt.Scan(&arr[i])
		i = i + 1
	}
	fmt.Print("Enter the number to be searched:")
	fmt.Scan(&num)
	index := linear_search(arr,n,num)
	fmt.Print("Element ",num," found at location ",index+1)
}
