package main

import (
	"fmt"
	"strings"
)

const (
	str = "kanguru"
	substr = "kang"
)

func main(){
	re := strings.Contains(str,substr)
	fmt.Println(re)
}

