package mapandstruct

import "fmt"

func Learn() {
	//Maps(key, value)
	// var m map[string]int
	m := map[string]int{"token": 10, "access": 20}
	fmt.Println(m)
	fmt.Println(m["access"])
	m["access"] = 200
	fmt.Println(m["access"])

	//loop
	for key, v := range m {
		fmt.Printf("%v => %v\n", key, v)
	}

	//delete map
	delete(m, "access")
	for key, v := range m {
		fmt.Printf("%v => %v\n", key, v)
	}

	//add map
	m["golang"] = 400
	for key, v := range m {
		fmt.Printf("%v => %v\n", key, v)
	}

	//struct

}
