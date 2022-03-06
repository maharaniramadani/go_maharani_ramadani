package main

import (
	"fmt"
)

type Student struct {
	name []string
	score []int
}

func(s Student) Avarage() float64{
	mapp :=make(map[string]int)

	for _,i:=range s.score{
		mapp[" "] += i
	}
	return float64(mapp[" "]/len(s.score))

}

func (s Student) Min() (min int, name string){
	mapp := map[string]int{
		" ": s.score[0],
	}

	var input int
	for i,value := range s.score{
		if mapp[" "] > value {
			mapp[" "] = value
			input = i
		}
	}
	return mapp[" "], s.name[input]

}

func (s Student) Max() (max int, name string){
	mapp := map[string]int{
		" ": s.score[0],
	}

	var input int
	for i,value := range s.score{
		if mapp[" "] < value {
			mapp[" "] = value
			input = i
		}
	}
	return mapp[" "], s.name[input]

	
}

func main(){
	var a = Student{}

	for i := 0; i<3;i++{
		var name string
		fmt.Print("input "+ string(i)+ " student name : ")
		fmt.Scan(&name)
		a.name=append(a.name, name)
		var score int
		fmt.Print("input " + name+ " score : ")
		fmt.Scan(&score)
		a.score=append(a.score, score)
	}
	fmt.Println("======================================")
	fmt.Println("\nnilai rata-rata: ", a.Avarage())
	scoremax,namemax :=a.Max()
	fmt.Println("nama : " +namemax+ " dengan nilai: " ,scoremax)
	scoremin,namemin :=a.Min()
	fmt.Println("nama : " +namemin+ " dengan nilai: " ,scoremin)
}