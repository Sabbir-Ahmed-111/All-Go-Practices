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
    fmt.Println("second",result)

	return
	
}

func cal() int{

	result :=0
	fmt.Println("first",result)

	show := func ()  {

		result = result + 10
        fmt.Println("defer",result)
		
	}
	
	defer show ()

	result = 5
    fmt.Println("second",result)

	return result
	
}

func main()  {

	a := calculator()
    fmt.Println("main first",a)

	b :=cal()
	fmt.Println("main second",b)


	
}