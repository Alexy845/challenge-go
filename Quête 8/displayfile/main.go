package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := ""
	if len(os.Args) == 2 {
		arg = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
	file, err := ioutil.ReadFile(arg)
	if err != nil {
		return
	}
	str := ""
	for _, i := range file {
		str += string(i)
	}
	fmt.Print(str)
}
