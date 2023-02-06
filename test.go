package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data_fale, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Не могу прочитать")
	}
	fmt.Println(string(data_fale))
}
