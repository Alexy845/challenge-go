package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func cat(r io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func main() {
	var output []byte
	var err error
	args := os.Args[1:]
	if len(args) > 0 {
		for _, file := range args {
			data, err := os.Open(file)
			if err != nil {
				for _, i := range "ERROR: open asd: no such file or directory\nexit status 1" {
					z01.PrintRune(i)
				}
				return
			}
			output, err = cat(data)
			if err != nil {
				return
			}
			for _, i := range string(output) {
				z01.PrintRune(i)
			}
		}
	} else {
		output, err = cat(os.Stdin)
		if err != nil {
			return
		}
		for _, i := range string(output) {
			z01.PrintRune(i)
		}
	}
}
