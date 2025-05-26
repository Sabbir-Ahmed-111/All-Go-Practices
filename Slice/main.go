package main

 import "fmt"

 func main(){

	arr:= [5]int8{1,2,3,4,5}
    // array theke slice
	s:=arr[1:4]//2,3,4
	fmt.Println(s)

    //slice theke slice
     s1:=s[1:2]//3
	 fmt.Printf("%d\n",s1)
	 fmt.Println(len(s1),cap(s1))

	 //make function slice

	 a:=make([]int,5,10)

	 a[1]=25
	 a[4]=35

	 //a = append(a,4,5,7,7,9,9) //ai vabe value add kora jabe
	 

	 fmt.Println(a)
	 fmt.Println(len(a))
	 fmt.Println(cap(a))

	//Empty Slice or nil slice 
	var x []int

	x = append(x, 1)// ptr-1 len-1, cap-1
	x = append(x, 2)//ptr-1,2 len-2, cap-2
	x = append(x, 3)//ptr-1,2,3 len-3, cap-6
	x = append(x, 4)//ptr-1,2,3,4 len-4, cap-6

	y := x

	x = append(x, 5)//prt-1,2,3,4,5 len-5, cap-6
	y = append(y, 6)//prt-1,2,3,4,6 len-5, cap-6
	x = append(x, 7)//prt-1,2,3,4,5,7 len-6, cap-6
	y = append(y, 10)//prt-1,2,3,4,6,10 len-6, cap-6

	x[2]=10 //prt-1,2,10,4,5,7 len-6, cap-6
	y[3]=100 //prt-1,2,3,100,6,10 len-6 cap-6

	fmt.Println(x)
    fmt.Println(y)



 }