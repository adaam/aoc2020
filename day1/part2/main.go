package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	var lines []int
	for _, v := range lines_str {
		str, _ := strconv.Atoi(v)
		lines = append(lines, str)
	}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			for k := j + 1; k < len(lines); k++ {
				if lines[i]+lines[j]+lines[k] == 2020 {
					fmt.Println(lines[i] * lines[j] * lines[k])
				}
			}
		}
	}
}
