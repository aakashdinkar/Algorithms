package main

import "fmt"

func circle_sort(arr[] int, start int,end int, swaps int)int{
	s := start
	e := end
	mid := (end - start)/2
	if start == end {
		return swaps
	}

	for(start < end){
		if arr[start] > arr[end] {
			arr[start],arr[end] = arr[end],arr[start]
			swaps = swaps + 1
		}
		start++
		end--
	}
	if start == end {
		if arr[start] > arr[end+1] {
			arr[start],arr[end+1] = arr[end+1],arr[start]
			swaps++
		}
	}
	swaps = circle_sort(arr,s,s+mid,swaps)
	swaps = circle_sort(arr,s+mid+1,e,swaps)
	return swaps
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
	for circle_sort(arr, 0, n-1, 0) != 0 {
    }
    fmt.Printf("Sorted  : %v\n\n", arr)
}
