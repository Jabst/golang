package main

import ("fmt")

func Sqrt(x float64) float64{
	z := float64(1)
	x1 := float64(0)
	for j := 0 ; j < 10 ; j++{
		x1 = z - (((z * z) - float64(x))/(2*z))
		z = x1
	}

	return x1
}

func main(){
	fmt.Println(Sqrt(2))
}