package main


import "fmt"



func value( arr *[5]int8, arr2 *[3]string){

	fmt.Printf("Integer numbers : %d\n",arr)
	fmt.Printf("String value :%s\n",arr2)

}

//custom data types

type Personal struct {
	Name string
	Age uint8
	Birthday int16
	Country string
	
}

func main(){
	

	arr := [5] int8{2,3,4,6,70}
	arr2 :=[3] string{"I","Love","You"}
	value(&arr,&arr2)
    //
	per := Personal{
       Name : " Sabbir",
	   Age: 25,
	   Birthday: 1999,
	   Country: "Bangladesh",

	}
   //fmt.Println(per)

   p := &per
   fmt.Println(*p)

}