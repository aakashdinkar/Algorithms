package main

	import "fmt"

	func naive_search(pat, text string){
		m:=len(pat)
		n:=len(text)
		for i:=0;i<n-m;i++{
			var j int
			for j=0;j<m;j++{
				if text[i+j] != pat[j]{
					break;
				}
			}
			if j == m{
				fmt.Println("Pattern found at Index ",i)
			}
		}
	}

	func main(){
		var pat,text string
		fmt.Print("Enter the text:")
		fmt.Scan(&text)
		fmt.Print("Enter the pattern:")
		fmt.Scan(&pat)
		naive_search(pat,text)
	}
