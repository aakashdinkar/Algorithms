package main

import "fmt"

func merge(a[] int, start int, end int, mid int){
	n1 := mid-start+1
	n2 := end-mid
	i,j,k := 0,0,0
	arr := make([]int,n1+n2)
	for(i < n1 && j < n2){
		if a[i] <= a[j]{
			arr[k] = a[i]
			k++
			i++

		}else{
			arr[k]=a[j]
			k++
			j++
		}
	}
	for(i<n1){
		arr[k] = a[i]
		k++
		i++
	}
	for(j<n2){
		arr[k]=a[j]
		k++
		j++
	}
	for i=0;i<n1+n2;i++{
		a[i]=arr[i]
	}
}

func merge_sort(arr[]int, start int, end int){
	if start > end{
		return 
	}
	mid := (start + end)/2
	mergesort(arr,start,end)
	mergesort(arr,mid+1,end)
	merge(arr,start,mid,end)
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
	merge_sort(arr,0,n)
	fmt.Printf("Sorted  : %v\n\n", arr)
}
