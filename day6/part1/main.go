package main

import (
	"fmt"
	"io/ioutil"

	//"strconv"
	"sort"
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

	var eva string
	sum := 0

	for _, v := range linesStr {
		if len(v) > 0 {
			eva = eva + v
		} else {
			evaSplit := strings.Split(eva, "")
			sort.Strings(evaSplit)
			var tmp string
			cnt := 0
			for k, v := range evaSplit {
				if k == 0 {
					tmp = v
					cnt++
				} else if v != tmp {
					tmp = v
					cnt++
				}
			}
			sum += cnt
			eva = ""
		}
	}
	fmt.Println(sum)
}
