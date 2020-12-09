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

func replace(str string, replaceTo0 string, replaceTo1 string) string {
	tmp1 := strings.ReplaceAll(str, replaceTo0, "0")
	result := strings.ReplaceAll(tmp1, replaceTo1, "1")
	return result
}

func main() {
	content, err := ioutil.ReadFile("input")
	check(err)
	lines_str := strings.Split(string(content), "\n")

	biggest := int64(0)
	for _, v := range lines_str {
		if len(v) <= 0 {
			break
		}
		row_bin := replace(v[0:7], "F", "B")
		col_bin := replace(v[7:10], "L", "R")
		row, _ := strconv.ParseInt(row_bin, 2, 8)
		col, _ := strconv.ParseInt(col_bin, 2, 4)
		if (row*8 + col) > biggest {
			biggest = row*8 + col
		}
	}
	fmt.Println("Biggest numbet is: " + strconv.Itoa(int(biggest)))
}
