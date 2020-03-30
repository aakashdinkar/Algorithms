package main

import "fmt"

func getmax(arr[] int,n int)int{
	max:=arr[0]
	for i:=0;i<n;i++{
		if arr[i] > max{
			max = arr[i]
		}
	}
	return max
}

func countsort(arr[] int, n,exp int){
	var max int 				 
 	max = getmax(arr,n)
 	i:=0
 	out:=make([]int, n)
 	count:=make([]int, max-0+1) 	
 	i=0
 	for(i<max){
 		count[i]=0
 		i++;
 	}

 	i=0
 	for(i<n){
 		count[(arr[i]/exp)%10]++
 		i++
 	}
 	i=1
 	for(i<max+1){
 		count[i]+=count[i-1]
 		i=i+1;
 	}
 	i=0
 	for(i<n){					
 		out[count[(arr[i]/exp)%10]-1] = arr[i]
 		count[(arr[i]/exp)%10]--
 		i++
 	}
	for i=0;i<n;i++{
		arr[i]=out[i]
	}
}

func radix_sort(arr[] int, n int){
	
	max := getmax(arr,n)

	for exp:=1;max/exp>0;exp*=10{
		countsort(arr,n,exp)
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
	radix_sort(arr,n)
	fmt.Printf("Sorted  : %v\n\n", arr)
}
