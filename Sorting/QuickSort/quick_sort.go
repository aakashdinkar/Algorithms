package main

import "fmt"

func partition(arr[] int, start,end int)int{
	pivot:= arr[end]
	i:= start-1
	for j:=start;j<end;j++{
		if arr[j] < pivot{
			i++
			arr[i],arr[j]=arr[j],arr[i]
		}
	}
	arr[i+1],arr[end]=arr[end],arr[i+1]
	return i+1
}

func quick_sort(arr[] int, start,end int){
	if start<end{
		p:=partition(arr,start,end)
		quick_sort(arr,start,p-1)
		quick_sort(arr,p+1,end)
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
	quick_sort(arr,0,n-1)
	fmt.Printf("Sorted  : %v\n\n", arr)
}
