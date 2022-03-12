package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var wg =sync.WaitGroup{}

func LetterFrequency(sentence string, i int, kapital bool){
	n := 97
	if kapital {
		n = 65
	}
	char := string(rune(i+n))
	letterFreq := strings.Count(sentence,char)
	if letterFreq != 0 {
		fmt.Printf("%s : %d \n",char,letterFreq)
	}
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(3)
	sentence :="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	for i := 0; i<26; i++{
		wg.Add(2)
		go LetterFrequency(sentence,i,false) 
		go LetterFrequency(sentence,i,true)  
	}
	wg.Wait()
}