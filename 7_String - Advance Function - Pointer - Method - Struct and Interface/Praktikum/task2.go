package main
import "fmt"

func caesar(offset int, input string)string{
	offset = (offset % 26 + 26) % 26
    new := make([]byte, len(input))
    for i := 0; i < len(input); i++ {
        t := input[i]
        var a int
        switch {
        case 'a' <= t && t <= 'z':
            a = 'a'
        default:
            new[i] = t
            continue
		}
        new[i] = byte(a + ((int(t)-a)+offset)%26)
    }
    return string(new)
}


func main(){
	fmt.Println(caesar(3,"abc"))
	fmt.Println(caesar(2,"alta"))
}
