package variables

import (
	"fmt"
	"strconv"
)

var fullname string //global
var email string = "pachara.prap@gmail.com"
var c, python bool = true, false

const vat int = 7

func Learn() {
	fullname = "Pachara"
	//fmt.Println(fullname)
	fmt.Printf("Hello %s \nEmail %s\n", fullname, email)
	fmt.Println(c, python)

	//must on function
	companyName := "Patchy"
	isShow := true
	age := 10
	fmt.Println(companyName, isShow, age)
	fmt.Printf("%T \n", isShow)

	f := float64(age)
	fmt.Printf("%T \n", f)

	//Convert to string
	s := strconv.Itoa(vat)
	fmt.Println("Vat is " + s)

}
