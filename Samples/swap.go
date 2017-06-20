package main

import "fmt"

func swap(x int, y string)(string, int){
	return y, x
}

func main(){
	a, b := swap(4,"ola")
	fmt.Println(a, b)
}