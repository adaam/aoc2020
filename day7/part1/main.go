package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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

	for _, v := range linesStr {
		r := regexp.MustCompile(`(?P<number>\d )?(?P<color>[A-z]+ [A-z]+) bags`)
		matches := r.FindAllStringSubmatch(v, -1)
		for k1, v1 := range matches {
			for k2, v2 := range v1 {
				fmt.Println(k2, v2, " v2\n")
			}
			fmt.Println(k1, v1, " v1\n")
		}
		fmt.Println(v, "v \n")
	}
	m := make(map[string]int)
	m["a b c"] = 1
	fmt.Println(m)
}
