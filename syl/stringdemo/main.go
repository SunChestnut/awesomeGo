package main

import (
	"fmt"
	"strings"
)

/**
*
* TODO
*
* @author  sun
* @date 2022/11/10 10:36
 */

func main() {
	s := "abcdefg"
	slice := strings.Split(s, "")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
