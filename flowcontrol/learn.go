package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func Learn() {
	//if
	score := 10
	if score == 10 {
		fmt.Println("very good!")

	} else {
		fmt.Println("try again!")
	}

	//switch
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s \n", os)
	}

	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}
	content := string(b)
	fmt.Println(content)

	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//break
	for j := 5; j >= 0; j-- {
		if j == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(j)
		time.Sleep(2 * time.Second)
	}
}
