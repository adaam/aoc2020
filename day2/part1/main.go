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
	var arr [][]string
	row_cnt := 0
	for k, v := range lines_str {
		arr = append(arr, strings.Split(v, " "))
		char_cnt := 0
		spec_char := strings.Split(arr[k][1], ":")[0]
		valid := strings.Split(arr[k][0], "-")
		valid_start, _ := strconv.Atoi(valid[0])
		valid_end, _ := strconv.Atoi(valid[1])
		for _, b := range arr[k][2] {
			if string(b) == spec_char {
				char_cnt++
			}
		}
		if char_cnt >= valid_start && char_cnt <= valid_end {
			row_cnt++
		}
	}
	fmt.Println("Correct rows: " + strconv.Itoa(row_cnt))
}
