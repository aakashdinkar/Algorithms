package main

import "fmt"

func binary_search(arr[] int, n int, num int) int {
	start,end := 0,n
	for(start <= end){
		mid := (start + end)/2
		if arr[mid] == num {
			return mid
		}else if arr[mid] < num {
			start = mid + 1
		}else{
			end = mid - 1
		}
	}
	return -1

}

func main(){
	var n,num int
	fmt.Print("Enter number of elements:")
	fmt.Scan(&n)
	fmt.Println("Enter the elements in sorted order")
	arr := make([]int, n)
	i:=0
	for(i<n){
		fmt.Print("Enter element ",i+1," :")
		fmt.Scan(&arr[i])
		i = i + 1
	}
	fmt.Print("Enter the number to be searched:")
	fmt.Scan(&num)
	index := binary_search(arr,n,num)
	fmt.Print("Element ",num," found at location ",index+1)
}
