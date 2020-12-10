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
	linesStr := strings.Split(string(content), "\n")
	linesStr[0] = ""

	charMap := make(map[int]map[string]int)

	for _, v := range linesStr {

		for k, ch := range v {
			//fmt.Println(k, string(ch))
			charMap[k] = map[string]int{}
			charMap[k][string(ch)] = 1
		}
		if len(v) == 0 {
			fmt.Println(len(charMap))
		}
	}

	//fmt.Println("1")
}
