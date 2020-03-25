package main

import "fmt"

func binary_search(arr[] int,l int, r int, num int) int {
	start,end := l,r
  if(start > end){
    return -1
  }
		mid := (start + end)/2
		if arr[mid] == num {
			return mid
		}else if arr[mid] < num {
			return binary_search(arr,mid+1,end, num)
		}else{
			return binary_search(arr,start,mid-1, num)
		}
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
	index := binary_search(arr,0,n,num)
	fmt.Print("Element ",num," found at location ",index+1)
}
