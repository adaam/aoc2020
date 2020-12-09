package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
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
	cnt := 0

	for _, v := range lines_str {
		cnt_field := 0
		if len(v) > 0 {
			eva = eva + v + " "
		} else {
			eva = strings.TrimSpace(eva)
			fields := strings.Split(strings.Replace(string(eva), " ", ":", -1), ":")
			
			for i := 0; i < len(fields) ; i += 2 {
				if i == 0 { 
					fmt.Println(fields, cnt_field, cnt)
				 }
				key := fields[i]
				value := fields[i+1]

				switch key {
				case "byr":
					int_value, _ := strconv.Atoi(value)
					if int_value >= 1920 && int_value<= 2002 {
						cnt_field++
					}
				case "iyr":
					int_value, _ := strconv.Atoi(value)
					if int_value >= 2010 && int_value<= 2020 {
						cnt_field++
					}
				case "eyr":
					int_value, _ := strconv.Atoi(value)
					if int_value >= 2020 && int_value<= 2030 {
						cnt_field++
					}
				case "hgt":
					hgt_length := len(fields[i+1])
					last_char := string(value[hgt_length-1])
					if last_char == "m" || last_char == "n" {
						int_value, _ := strconv.Atoi(value[0:hgt_length-2])
						if last_char == "m" && int_value >= 150 && int_value <= 193 {
							cnt_field++
						} else if last_char == "n" && int_value >= 59 && int_value <= 76 {
							cnt_field++
						}
					}
				case "hcl":
					format, _ := regexp.Compile("#[0-9a-f]{6}")
					if format.MatchString(value) {
						cnt_field++
					}
				case "ecl":
					list := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
					for _, v := range list {
						if value == v {
							cnt_field++
						}
					} 
				case "pid":
					format, _ := regexp.Compile("[0-9]{9}")
					if format.MatchString(value) {
						cnt_field++
					}
				}
				fmt.Println(fields ,key, value, cnt_field)
				//fmt.Println(cnt_field)
			}
			if cnt_field == 7 {
				cnt++
			}

			eva = ""

		}
	}
	fmt.Println(cnt)
}