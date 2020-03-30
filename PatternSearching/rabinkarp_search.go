package main

import "fmt"

func rabinkarp_search(pat,text string, q int){
	m:=len(pat)
	n:=len(text)
	var i,j int
	p,t,h,d:=0,0,1,256

	for i=0;i<m-1;i++{
		h = (h*d)%q
	}
	for i=0;i<m;i++{
		p=(d*p+int(pat[i]))%q
		t=(d*t+int(pat[i]))%q
	}
	for i=0;i<=n-m;i++{
		if p==t{
			for j=0;j<m;j++{
				if text[i+j] != pat[j]{
					break
				}
			}
			if j==m{
				fmt.Println("Pattern found at index ",i)
			}
		}
		if i<n-m{
			t=(d*(t-int(text[i])*h)+int(text[i+m]))%q
			if t<0{
				t=t+q
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
	rabinkarp_search(pat,text,101)
}
