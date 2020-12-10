package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"strings"
	"sort"
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
			eva_split := strings.Split(eva, "")
			sort.Strings(eva_split)
			var tmp string
			cnt := 0
			for k, v := range eva_split {
				if k == 1 {
					tmp = v
					cnt++
				} else if v != tmp {
					cnt++
				}
			}
			fmt.Println(eva_split)
			fmt.Println(cnt)
			eva = ""
		}
	}
}