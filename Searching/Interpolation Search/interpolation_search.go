package main

import "fmt"

func interpolation_search(arr []int, n int, num int) int {

    min, max := arr[0], arr[n-1]

    start, end := 0, n-1

    for {
        if num < min {
            return start
        }

        if num > max {
            return end + 1
        }
        var guess int
        if end == start {
            guess = end
        } else {
            size := end - start
            pos := int(float64(size-1) * (float64(num-min) / float64(max-min)))
            guess = start + pos
        }
        if arr[guess] == num {
            for guess > 0 && arr[guess-1] == num {
                guess--
            }
            return guess
        }
        if arr[guess] > num {
            end = guess - 1
            max = arr[end]
        } else {
            start = guess + 1
            min = arr[start]
        }
    }
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
	index := interpolation_search(arr,n,num)
	fmt.Print("Element ",num," found at location ",index+1)
}
