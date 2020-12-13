package main

import (
	"fmt"
	"io/ioutil"
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

	charCount := make(map[string]int)
	var s []map[string]int
	total := 0

	for _, v := range linesStr {
		for _, ch := range v {
			//fmt.Println(k, string(ch))
			charCount[string(ch)]++
		}

		if len(v) > 0 {
			s = append(s, charCount)
		}

		charCount = make(map[string]int)

		if len(v) == 0 {
			//fmt.Println(s[0])
			cnt := 0
			ok := true
			for chVerify, _ := range s[0] {
				//fmt.Println(chVerify, "\n")
				for _, charSet := range s {
					_, ok := charSet[chVerify]
					if !ok {
						fmt.Println("wrong: ", chVerify)
						break
					} else {
						//fmt.Println(charSet, len(s))
						cnt++
					}
				}
				if !ok {
					fmt.Printf("%v", ok)
					break
				}
				if cnt == len(s) {

					total += cnt
					fmt.Println("total: ")
					fmt.Println(cnt, len(s), total)
				}
			}
			fmt.Println(total)
			s = s[:0]
		}
	}

	//fmt.Println("1")
}
