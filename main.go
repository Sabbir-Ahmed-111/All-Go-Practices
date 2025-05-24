package main

import (
	"fmt"
	"hello/information"

)


 

func main(){

	fmt.Println("Custom Package")
	information.Details(10,3,information.Call)

	//function return
   sum:=information.Print()

   sum(20,30)

}