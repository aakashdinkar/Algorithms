package main

import "fmt"

func count_Sort(arr[] int,out[] int, n int){
  var max int 				 
 	max = arr[0]
 	i:=0
 	for(i<n){
 		if arr[i] > max{
 			max=arr[i]

 		}
 		i=i+1
 	}
 	count:=make([]int, max-0+1) 	
 	i=0
 	for(i<max){
 		count[i]=0
 		i++;
 	}

 	i=0
 	for(i<n){
 		count[arr[i]]++
 		i++
 	}
 	i=1
 	for(i<max+1){
 		count[i]+=count[i-1]
 		i=i+1;
 	}
 	i=0
 	for(i<n){					
 		out[count[arr[i]]-1] = arr[i]
 		count[arr[i]]--
 		i++
 	}
}

func main(){
	var n int
	fmt.Print("Enter number of elements:")
	fmt.Scan(&n)
	fmt.Println("Enter the elements of array")
	arr := make([]int, n)
	out:=make([]int, n)				
	i:=0
	for(i<n){
		fmt.Print("Enter element ",i+1," :")
		fmt.Scan(&arr[i])
		i = i + 1
	}
	count_Sort(arr,out,n)
	fmt.Println("Sorted Array is : ",out)
}
