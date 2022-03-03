// package main

// import (
// 	"fmt"

// )

//  func prime(num int)bool{
//    var i int
//     for (i := 2; i * i <= num; i++){
//       if (num % i === 0)
//           return false;
//     }
//     return num > 1;
// }

//  func main(){
// 	 fmt.Println(prime(1000000007))
//  }

// package main

// import (
//     "fmt"
// )

// func main() {

// }

//         s1 := []string{"James", "Wagner", "Christene", "Mike"}
//         s2 := []string{"Paul", "Haaland", "Patrick"}

//         s3 := append(s1, s2...)
//         fmt.Println(s3) // [James Wagner Christene Mike Paul Haaland Patrick]

// a := []int{1, 2, 3}
// b := []int{3, 5, 6}
// fmt.Println(a)
// fmt.Println(b)
// fmt.Println()

// c := append(a, b)  // cannot use b (type []int) as type int in append
// c := append(a, b...)
// fmt.Println(c)
// fmt.Printf("%T\n", c)
// fmt.Println(len(c))
// fmt.Println(cap(c))
// }

// package main

// import (
//   "fmt"
// )

// func ArrayMerge(arrayA, ArrayB []string) []string {
//   newArray := arrayA
//   for index := range ArrayB {
//     newArray = append(newArray, ArrayB[index])
//   }
//   return newArray
// }

// func ArrayMerge(nums1 []string, m int, nums2 []string, n int)  {
//     i := len(nums1) - 1
//     for n > 0 {
//         if m == 0 || nums1[m - 1] <= nums2[n - 1] {
//             nums1[i] = nums2[n - 1]
//             i--
//             n--
//         } else {
//             nums1[i] = nums1[m - 1]
//             i--
//             m--
//         }
//     }

// }

// func main() {
//   // Test cases
//   fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
//   // ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

//   fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
//   // ["sergei", "jin", "steve", "bryan"]

//   fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
//   // ["alisa", "yoshimitsu", "devil jin", "law"]

//   fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
//   // ["devil jin", "sergei"]

//   fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
//   // ["hwoarang"]

//   fmt.Println(ArrayMerge([]string{}, []string{}))
//   // []
// }

// func linear(n int) {
// 	for i := 1; i <= n; i++ { // execution time increases according to variable n
// 		fmt.Println("loop execution:", i)
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

// start := time.Now()
// linear(5)
// fmt.Println("time elapsed:", time.Since(start))
// start = time.Now()
// linear(10)
// fmt.Println("time elapsed:", time.Since(start))

// time complecity lebih cepat

// package main

// import (
// 		"fmt"
// 		"math/big"
// 		"time"
// 	)

// 	// func Pangkat(base, pangkat int) int {
// 	// 	var i, result int
// 	// 	result = 1
// 	// 	for i = 0; i < pangkat; i++ {
// 	// 			result *= base
// 	// 			time.Sleep(1000 * time.Millisecond)
// 	// 	}
// 	// 	return result
// 	// }

// 	// func main() {
// 	// 	start := time.Now()
// 	// 	fmt.Println(Pangkat(2, 3))
// 	// 	fmt.Println("time elapsed:", time.Since(start))
// 	// 	fmt.Println(Pangkat(5, 3))
// 	// 	fmt.Println("time elapsed:", time.Since(start))
// 	// 	fmt.Println(Pangkat(10, 2))
// 	// 	fmt.Println("time elapsed:", time.Since(start))
// 	// 	fmt.Println(Pangkat(2, 5))
// 	// 	fmt.Println("time elapsed:", time.Since(start))
// 	// 	fmt.Println(Pangkat(7, 3))
// 	// 	fmt.Println("time elapsed:", time.Since(start))
// 	// }

// 	// func main() {
// 	// 	const n =1000000007
// 	// 	if big.NewInt(n).ProbablyPrime(0) {
// 	// 		fmt.Println(n, " adalah bilangan prima")
// 	// 		time.Sleep(100 * time.Millisecond)
// 	// 		start := time.Now()
// 	// 		fmt.Println("lama waktu : ", time.Since(start))
// 	// 	} else {
// 	// 		fmt.Println(n, " adalah  bukan  bilangan prima")
// 	// 		time.Sleep(100 * time.Millisecond)
// 	// 		start := time.Now()
// 	// 		fmt.Println("lama waktu : ", time.Since(start))
// 	// 	}
// 	// }
// 	func angkaprima (num int64){
// 	// prima := true
// 	// if num == 0 || num == 1 {
// 	// 	fmt.Printf(" %d adalah bukan bilangan prima\n", num)
// 	// 	time.Sleep(200 * time.Millisecond)
// 	// }else{
// 	// 	for i := 2; i <= num/2; i++ {
// 			if big.NewInt(num).ProbablyPrime(0) {
// 				fmt.Println(num, "is prime")
// 				time.Sleep(1000 * time.Millisecond)
// 			// start := time.Now()
// 			// fmt.Println("time elapsed:", time.Since(start))
// 		} else {
// 			fmt.Println(num, "is not prime")
// 			time.Sleep(1000 * time.Millisecond)
// 			// start := time.Now()
// 			// fmt.Println("time elapsed:", time.Since(start))
// 		}
// 	}

// 	func main(){
// 		start := time.Now()
// 		angkaprima(1000000007)
// 		fmt.Println("time elapsed:", time.Since(start))
// 		// // start := time.Now()
// 		// angkaprima(13)
// 		// fmt.Println("time elapsed:", time.Since(start))
// 		// // start := time.Now()
// 		// angkaprima(17)
// 		// fmt.Println("time elapsed:", time.Since(start))
// 	}

// //task5
// package main

// import "fmt"

// /*--
// 	Solution using hashing
// --*/

// /* --
// 	time complexity is O(n)
// --*/

// // Function to find a pair in an array with a given sum using hashing
// func findPair(arr []int, n, sum int) {
// 	// create an empty map
// 	mp := make(map[int]int)

// 	// do for each element of array
// 	for i := 0; i < n; i++ {
// 		// add elem to the map
// 		mp[arr[i]] = i
// 		// check if pair (arr[i], sum-arr[i]) exists
// 		// if the difference is seen before, print the pair
// 		elem1, ok1 := mp[arr[i]]
// 		elem2, ok2 := mp[sum-arr[i]]
// 		//fmt.Println(ok1, ok2)
// 		if ok1 && ok2 {
// 			fmt.Printf("Pair found at: %d, %d", elem1, elem2)
// 			return
// 		}
// 	}
// 	// No pair with the given sum exists in the array
// 	fmt.Println("Pair not found")
// }

// func main() {
// 	arr := []int{8, 7, 2, 5, 3, 1}

// 	// call the function
// 	findPair(arr, len(arr), 10)
// }

// package main

// import (
// 	"fmt"
// 	// "math"
// 	"time"
// )

//  func prime(num int)bool{
// 	for i := 2; i <= int(num)/2; i++ {
//         if num%i == 0 {
//             return false
// 			time.Sleep(1000 * time.Millisecond)
//         }
//     }
//     return num > 1
//  }

//  func main(){
// 	start := time.Now()
// 	 fmt.Println(prime(1000000007))
// 	 fmt.Println("lama waktu:", time.Since(start))
//  }

package main 
import "fmt"

func luas (jari float32)float32{
	luas:=3.14*jari*jari
	return luas
}

func luassel(jari,tinggi float32)float32{
	ls:=2*3.14*jari*tinggi
	return ls
}

func main (){
	var jari float32 =7.0
	var tinggi float32 =10.0
	
}
// func Compare(a, b string) int {
// 	if a == b {
// 		return 0
// 	}
// 	if a < b {
// 		return -1
// 	}
// 	return +1
// }
// func LCS(s1, s2 string) string {
// 	var m = make([][]int, 1 + len(s1))
// 	for i := 0; i < len(m); i++ {
// 	  m[i] = make([]int, 1 + len(s2))
// 	}
// 	longest := 0
// 	x_longest := 0
// 	for x := 1; x < 1 + len(s1); x++ {
// 	  for y := 1; y < 1 + len(s2); y++ {
// 		if s1[x-1] == s2[y-1] {
// 		  m[x][y] = m[x-1][y-1] + 1
// 		  if m[x][y] > longest {
// 			longest = m[x][y]
// 			x_longest = x
// 		  }
// 		}
// 	  }
// 	}
// 	return s1[x_longest - longest: x_longest]
//   }
// fmt.Println(Compare("aka","akasi"))
// 	fmt.Println(Compare("kangguru","kang"))