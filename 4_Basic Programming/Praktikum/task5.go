
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
}
// package main

// import "fmt"

// func isPalindrome(input string) bool {
//   for i := 0; i < len(input)/2; i++ {
//     if input[i] != input[len(input)-i-1] {
//       return false
//     }
// }
// return true
// }

// func main() {
//   var str string
//   fmt.Scan(&str)
//   result := isPalindrome(str)
//   if result == true {
//     fmt.Println("palindrome")
//     } else {
//       fmt.Println("not palindrome")
//     }

// }