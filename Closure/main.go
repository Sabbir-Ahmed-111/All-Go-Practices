package main

import "fmt"

const p = 20
var x = 100


func add() func(){
	money := 50
	fmt.Println(x)



	show:=func(){
		money = money + p + x

		fmt.Println(money)
    
	}

	return show
	
}

func call(){
	sum := add()
	sum ()
	sum ()
}

func main()  {

	call()
	
}