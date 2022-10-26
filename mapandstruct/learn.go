package mapandstruct

import (
	"fmt"
)

func Learn() {
	//Maps(key, value)
	// var m map[string]int
	// var m2 = make(map[string]int)
	// m2["a"] = 10
	// m2["b"] = 20
	// fmt.Println(m2["b"])

	// m := map[string]int{"token": 10, "access": 20}
	// fmt.Println(m)
	// fmt.Println(m["access"])
	// m["access"] = 200
	// fmt.Println(m["access"])

	// //loop
	// for key, v := range m {
	// 	fmt.Printf("%v => %v\n", key, v)
	// }

	// //delete map
	// delete(m, "access")
	// for key, v := range m {
	// 	fmt.Printf("%v => %v\n", key, v)
	// }

	// //add map
	// m["golang"] = 400
	// for key, v := range m {
	// 	fmt.Printf("%v => %v\n", key, v)
	// }

	//struct
	type User struct {
		id       int
		username string
		password string
	}

	john := User{
		id:       1,
		username: "John Doe",
		password: "1234",
	}

	fmt.Println(john.username)
	john.password = "12345678"
	fmt.Println(john.password)

	user := []User{
		{id: 1, username: "Mary", password: "1234"},
		{id: 2, username: "Bob", password: "4567"},
	}

	fmt.Println(user[0])
	fmt.Println(user[1].username)

}
