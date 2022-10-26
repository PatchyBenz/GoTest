package pointers

import "fmt"

func Learn() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x) //$ = get memory address

	y := x
	fmt.Println(y)
	fmt.Println(&y)

	fmt.Println("--------------")
	//var p *int = &x
	p := &x //create pointer p int
	fmt.Println(p)
	fmt.Println(*p) // read value x to p pointer(dereference)

	fmt.Println("--------------")
	*p = 20
	fmt.Println(*p)
	fmt.Println(x)

	b := student{name: "Bob"}
	fmt.Println(b)
	setName(&b)
	fmt.Println(b)
}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Boy"
}

/*
	& if this operator push front of valuable will read memory address
	* if this operator push front of valuable will get value from memory address
	* if this operator push front of Type such as *student, *int .This keep value of memory address
*/
