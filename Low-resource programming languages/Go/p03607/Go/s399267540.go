
package main

import(
	. "fmt"
)

func main(){
	N:=0
	AL:=0
	le:=0
	Scanf("%d",&N)
	var A []int
  	Scanf("%d",&AL)
  	A=append(A,AL)

	for i:=1;i<N;i++{
		Scanf("%d",&AL)
      	
		for p:=0;p<len(A);p++{
			if AL==A[p]{
				A[p]=-1
				break
			}else if p+1==len(A){
				A=append(A,AL)
              	break
			}

		}

	}
	le=len(A)
	for p:=0;p<len(A);p++{
		if A[p]==-1{
			le--;
		}
	}
	Println(le)

}