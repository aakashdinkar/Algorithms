package main

import "fmt"

func kmp_search(pat,text string){
	m:=len(pat)
	n:=len(text)
	long:=make([]int, m)
	computelongest(pat,m,long)
	i,j := 0,0
	for(i<n){
		if pat[j] == text[i]{
			j++
			i++
		}
		if j==m{
			fmt.Print("Pattern found at index ",i-j)
			j = long[j-1]
		}else if i<n && pat[j]!=text[i]{
			if j!=0{
				j=long[j-1]
			}else {
				i++
			}
		}
	}
}

func computelongest(pat string, n int, long[] int){
	len:=0
	long[0]=0
	i:=1
	for(i<n){
		if pat[i] == pat[len]{
			len++
			long[i]=len
			i++
		}else{
			if len != 0{
				len = long[len-1]
			}else {
				long[i]=0
				i++
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
	kmp_search(pat,text)
}
