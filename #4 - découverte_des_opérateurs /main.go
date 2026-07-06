package main

import "fmt"


func main()  {
	
	var (
		x int
		y int
	)


	x = 7
	y = 3 

//+ - / * %
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(x % y)

fmt.Println("=================")

	//== != < <= > >= 
	fmt.Println(x == y) 
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println(x > y)
	fmt.Println(x >= y)

	fmt.Println("=================")
	//&& || 
	fmt.Println(x == y && x!=y) // et logique
	fmt.Println(x == y || x!=y) //ou logique

}


/*


+ - / * %     ===> Arithmetique 

== != < <= < <=    =====> Relationnel

&& ||     ======> Logique


*/
