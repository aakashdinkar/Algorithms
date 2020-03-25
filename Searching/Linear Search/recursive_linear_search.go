package main

import "fmt"

func linear_search(arr[] int, start int, end int, num int) int {
	if end < start{
		return -1
	}
	if arr[start] == num{
		return start
	}
	if arr[end] == num{
		return end
	}
	return linear_search(arr,start+1,end-1,num)
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
	index := linear_search(arr,0,n-1,num)
	fmt.Print("Element ",num," found at location ",index+1)
}
