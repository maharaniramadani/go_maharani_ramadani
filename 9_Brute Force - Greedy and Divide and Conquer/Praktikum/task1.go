package main
import "fmt"


func simpEq(a,b,c int) {
	var pjg int
	if c <a || c < b{
		fmt.Println("no solution")
		return

	}
	if a >b{
		pjg=a
	}else {
		pjg=b
	}
	if pjg <=5{
		fmt.Println("no solution")
		return

	}

	for x := 1; x <=(pjg/2)-2 ;x++ {
		if x+(x+1)+(x+2)== a && x*(x+1)*(x+2)== b && (x*x)+((x+1)*(x+1))+((x+2)*(x+2))== c {
			fmt.Println(x,x+1,x+2)
			return
		}
		for y := x+1; y <= (pjg/2)-1; y++ {
			if x+y+(y+1)== a && x*y*(y+1)== b && (x*x)+(y*y)+((y+1)*(y+1))== c {
				fmt.Println(x,y,y+1)
				return
			}
			for z := y+1; z <= (pjg/2); z++ {
				if x+(y)+(z)== a && x*(y)*(z) ==b && (x*x)+((y)*(y))+((z)*(z)) ==c {
					fmt.Println(x,y,z)
					return

				}
			}
		}
	}
	fmt.Println("no solution")
}

func main() {
	simpEq(1,2,3)
	simpEq(6,6,14)   
}