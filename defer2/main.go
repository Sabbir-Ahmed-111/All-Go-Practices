package main

import"fmt"
 

func calculator() (result int) {

	fmt.Println("first",result)


	show := func ()  {

		result = result + 10
        fmt.Println("defer",result)
		
	}
	defer show ()


	result = 5
    
	p:=func (a int)  {
		fmt.Println(a)
	}
	defer p (result)

	defer fmt.Println(result)

    fmt.Println("second",result)

	defer fmt.Println(5)

	return
	
}

func main()  {
	a := calculator()
	fmt.Println(a)
	
}