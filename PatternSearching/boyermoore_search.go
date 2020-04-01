package main

import "fmt"

func bcharheuristic(pat string, m int, bchar[] int){
	var i int
	for i=0;i<256;i++{
		bchar[i]=-1
	}
	for i=0;i<m;i++{
		bchar[int(pat[i])]=i
	}
}

func max(a,b int)int{
	if a > b{
		return a
	}else {return b}
}

func search(pat,text string){
	m,n:=len(pat),len(text)

	bchar:=make([]int, 256)

	bcharheuristic(pat,m,bchar)

	s:=0
	for(s<=(n-m)){
		i:= m-1
		for(i>=0 && pat[i] == text[s+i]){
			i--
		}
		if i < 0{
			fmt.Print("Pattern found at index ",s)
			if s+m<n{
				s+=m-bchar[int(text[s+m])]
			}else{
				s+=1
			}
		}else {
			s+=max(1,i-bchar[int(text[s+i])])
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
