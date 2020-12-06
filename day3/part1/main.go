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
	forest_length := len(lines_str)
	//forest_width := len(lines_str[0]) + 1
	right_step := 3
	down_step := 1
	trees_cnt := 0
	for i, j := down_step, right_step; i < forest_length; i, j = i+down_step, j+right_step {
		if string(lines_str[i][j%31]) == "#" {
			trees_cnt++
		}
	}
	fmt.Println("Trees counts is: " + strconv.Itoa(trees_cnt))

}
