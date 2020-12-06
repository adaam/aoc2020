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
		//char_cnt := 0
		spec_char := string(strings.Split(arr[k][1], ":")[0])
		valid := strings.Split(arr[k][0], "-")
		valid_pos1, _ := strconv.Atoi(valid[0])
		valid_pos2, _ := strconv.Atoi(valid[1])
		pos1_char := string(arr[k][2][valid_pos1-1])
		pos2_char := string(arr[k][2][valid_pos2-1])

		if !(pos1_char == spec_char && pos2_char == spec_char) && (pos1_char == spec_char || pos2_char == spec_char) {
			row_cnt++
		}
	}
	fmt.Println("Correct rows: " + strconv.Itoa(row_cnt))
}
