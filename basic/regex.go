package main

import (
	"fmt"
	"regexp"

)

const text = `my email is leizhang.gamil@gmail.com 13094@qq.com`

func main() {
	re := regexp.MustCompile(`([a-z0-9A-Z_.]+)@([a-zA-Z0-9.]+)(\.[a-z.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Printf("%s\n", m)
	}
}
