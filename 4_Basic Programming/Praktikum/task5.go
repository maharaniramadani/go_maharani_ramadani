
// polindrom
package main

import "fmt"

func palindrome(input string) bool {
  for i := 0; i< len(input)/2; i++ {
    if input[i] != input[len(input)-i-1]{
      return false
    }
  }
	return true
}

func main() {
  fmt.Println(palindrome("katak"))
  fmt.Println(palindrome("civic"))
  fmt.Println(palindrome("kasur rusak"))
  fmt.Println(palindrome("mistar"))
  fmt.Println(palindrome("lion"))
<<<<<<< HEAD
}
=======
}
>>>>>>> 9a4e23a6d2bf8a32b86023b6fbb9d6ae57cd3a8d
