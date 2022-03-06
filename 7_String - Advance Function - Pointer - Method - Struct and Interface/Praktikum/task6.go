package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	nameEncode string
	score int
}

type cipher interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string{
	var nameEncode=""

	encod:=func(offset rune) rune {
		i := int(offset)+3
		if i >'z'{
			return rune(i-26)
		}else if i <'a'{
			return rune(i+26)
		}
		return rune(i)
	}
	nameEncode = strings.Map(encod, s.name)
	return nameEncode
}

func (s *student) Decode() string{
	var nameDecode=""

	dencod:=func(offset rune) rune {
		i := int(offset)-3
		if i >'z'{
			return rune(i-26)
		}else if i <'a'{
			return rune(i+26)
		}
		return rune(i)
	}
	nameDecode = strings.Map(dencod, s.nameEncode)
	return nameDecode
}

func main(){
	var menu int
	var a =student{}
	var c cipher=&a

	fmt.Print("[1]Encrypt \n[2]Decrypt \nchoose your menu: ")
	fmt.Scan(&menu)

	if menu == 1{
		fmt.Print("\ninput student name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nencode of student name " +a.name+ " is: " +c.Encode())
	}else if menu == 2{
		fmt.Print("\ninput student encode name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student name " +a.nameEncode+ " is: " +c.Decode())
	}else{
		fmt.Println("wrong input name menu")
	}
}


