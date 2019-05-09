package main

/*
	@author	Boosh
	@desc	Program to generate some pixel map of the mandelbrot set
*/

import (
	"fmt"
	"math"
)

//type complex
//contains a complex number and data about it as it applies to the mandelbrot set - 
//	it's size, and the # of iterations it took to stabilize or fly off to infinity
type complex struct {
	real	float64
	imag	float64
	size	float64
	count	int
}

func printComplex(c complex){
	if c.imag < 0 {
		fmt.Println(fmt.Sprintf("%f - %fi --- size: %f", c.real,-c.imag,getComplexSize(c)))
	} else {
		fmt.Println(fmt.Sprintf("%f + %fi --- size: %f", c.real,c.imag,getComplexSize(c)))
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
	return math.Sqrt((c.real*c.real)+(c.imag*c.imag))
}

func main(){
	fmt.Println("Howdy")

	//edge size of square for mandelbrot image, and mandelbrot image
	size := 200
	var image [200][200] complex
	step := 2.5/float64(size)

	//starting c = 1
	c := complex{real:0.0,imag:0.0}

	//number of count iterations for each pixel
	maxCount := 100

	//for each space, calculate if sizeof(z**2+c) <= 2, 
	for x,a := 0,-1.25;x<size;x++{
		for y,b := 0,-2.0;y<size;y++{
			image[x][y] = complex{real:a,imag:b}

			for count := 0;count < maxCount;count++{
				image[x][y] = zSquaredPlusC(image[x][y],c)
				image[x][y].size = getComplexSize(image[x][y])
				image[x][y].count = count
				if image[x][y].size > 2{break}
			}

			if image[x][y].size < 2{
				fmt.Println(fmt.Sprintf("Array [%d][%d], Coords: (%f,%f)",x,y,a,b))
			}

//			fmt.Println(fmt.Sprintf("At end of inner y loop (%d,%d), image: %f,%f | %f @ %d",x,y,image[x][y].real,image[x][y].imag,image[x][y].size,image[x][y].count))

			//increment real component
			b += step
		}
		//increment imaginary component
		a += step
	}
	return
	for x := -(size/2);x < (size/2);x++ {
		for y := 0;y < size;y++ {
			image[-x][y] = complex{real:(float64(x)/float64(size)),imag:(float64(y)/float64(size))}
			fmt.Println("Now analyzing point (%f,%f)",x,y)
			printComplex(image[-x][y])
			//analyze wether or not this square is in the mandelbrot set or not
			for count := 0;count<100;count++{
				image[-x][y] = zSquaredPlusC(image[-x][y],c)
				image[-x][y].size = getComplexSize(image[-x][y])
				fmt.Println(fmt.Sprintf("(%d,%d): current z: %f",x,y,image[x][y].size))
				image[-x][y].count = count
				if image[-x][y].size > 2 {
					fmt.Println("this pixel is outside the set:")
//					printComplex(image[x][y])
					break
				}
			}
			if image[-x][y].size < 2 {
				fmt.Println(fmt.Sprintf("Pixel created (%d,%d):",x,y))
				printComplex(image[-x][y])
			}
//			fmt.Println(fmt.Sprintf("Count after loop: %d",image[x][y].count))
		}

//				return
	}
}
