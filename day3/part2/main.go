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
	forest_width := len(lines_str[0])
	steps := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	mul_num := 1
	for _, v := range steps {
		right_step := v[0]
		down_step := v[1]
		trees_cnt := 0
		for i, j := down_step, right_step; i < forest_length; i, j = i+down_step, j+right_step {
			if string(lines_str[i][j%forest_width]) == "#" {
				trees_cnt++
			}
		}
		mul_num *= trees_cnt
	}
	fmt.Println("Multiplie all numbers result: " + strconv.Itoa(mul_num))
}
