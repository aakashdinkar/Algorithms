package main

import ("fmt"
		"math")

func jump_search(arr[] int, n int, num int) int {
	jump:=math.Sqrt(float64(n))			
	i:=0
	for(i<n){
		if arr[i] == num {
			return i
		} else if arr[i] > num {
			for(arr[i] != num){
				i=i-1
			}
      		return i	
		}
		i=i+int(jump)
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
	index := jump_search(arr,n,num)
  if index == -1{
      fmt.Print("Element not found")
  }
	fmt.Print("Element ",num," found at location ",index+1)
}
