package main

import "fmt"

func search(pat,text string){
	str:= pat + "$" + text
	l:=len(str)
	z:=make([]int, l)
	getZarr(str,z)
	for i := 0; i < l; i++ {
		if z[i] == len(pat){
			fmt.Println("Pattern found at index ",i-len(pat)-1)
		}
	}
}

func getZarr(str string, z[] int){
	n := len(str)
	var L,R,k int
	L,R = 0,0
	for i := 0; i < n; i++ {
		if i > R{
			L,R=i,i
			for(R<n && str[R-L] == str[R]){
				R++
			}
			z[i] = R-L
			R--
		}else{
			k = i-L
			if z[k] < R-i+1{
				z[i] = z[k]
			}else{
				L=i
				for(R<n && str[R-L] == str[R]){
					R++
				}
				z[i] = R-L
				R--
			}
		}
	}
}

func main(){
	var pat,text string
	fmt.Print("Enter the text:")
	fmt.Scan(&text)
	fmt.Print("Enter the pattern:")
	fmt.Scan(&pat)
	search(pat,text)
}
