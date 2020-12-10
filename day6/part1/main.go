package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := ioutil.ReadFile("input")
	check(err)
	lines_str := strings.Split(string(content), "\n")

	var eva string

	for _, v := range lines_str {
		if len(v) > 0 {
			eva = eva + v
		} else {
			fmt.Println(eva + "\n")
			eva = ""
		}
	}
}