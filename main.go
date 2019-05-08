package main

/*
	@author	Boosh
	@desc	Program to generate some pixel map of the mandelbrot set
*/

import (
	"fmt"
	"math"
)

type complex struct {
	real	int
	imag	int
}

func printComplex(c complex){
	if c.imag < 0 {
		fmt.Println(fmt.Sprintf("%d - %di --- size: %f", c.real,-c.imag,getComplexSize(c)))
	} else {
		fmt.Println(fmt.Sprintf("%d + %di --- size: %f", c.real,c.imag,getComplexSize(c)))
	}
}

func zSquaredPlusC(z complex, c complex) complex {
	var r complex

	//z^2
	//real component is real^2 - imag^2, imaginary component is 2*real*imag
	r.real = (z.real*z.real)-(z.imag*z.imag)
	r.imag = 2*z.real*z.imag

	//+c
	r.real = r.real+c.real
	r.imag = r.imag+c.imag

	return r
}

func getComplexSize(c complex) float64 {
	//the float64 is due to golang's math.Sqrt requiring float64 as input, not int
	return math.Sqrt(float64(c.real*c.real)+float64(c.imag*c.imag))
}

func main(){
	fmt.Println("Howdy")

	c := complex{real:1,imag:1}
	z := complex{real:0,imag:0}

	for i := 0;i<10;i++{
		z = zSquaredPlusC(z,c)
		printComplex(z)
	}
}
