package main

/*
	@author	Boosh
	@desc	Program to generate some pixel map of the mandelbrot set
*/

import (
	"fmt"
	"math"
//	"strings"
	"image"
	"image/color"
	"image/png"
	"os"
)

//type complex
//contains a complex number and data about it as it applies to the mandelbrot set - 
//	it's size, and the # of iterations it took to stabilize or fly off to infinity
type complex struct {
	real	float64	//real component
	imag	float64	//number coorespondin to imaginary component
	size	float64	//size of coordinate -- sqrt(real**2 + imag**2)
	count	int	//number of iterations taken to pass 2
	enab	byte	//is the pixel enabled (in the set)
	color	int	//0-255
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
	var imageArr [200][200] complex
	imgOut := image.NewNRGBA(image.Rect(0,0,200,200))

	step := 4.0/float64(size)

	//starting c = 1
	c := complex{real:-1.2,imag:0.1}

	//number of count iterations for each pixel
	maxCount := 1000

	//for each space, calculate if sizeof(z**2+c) <= 2, 
	for x,a := 0,-2.0;x<size;x++{
		for y,b := 0,-2.0;y<size;y++{
			imageArr[x][y] = complex{real:a,imag:b}

			for count := 0;count < maxCount;count++{
				imageArr[x][y] = zSquaredPlusC(imageArr[x][y],c)
				imageArr[x][y].size = getComplexSize(imageArr[x][y])
				imageArr[x][y].count = count
				if imageArr[x][y].size > 2.0 {break}
			}

			if imageArr[x][y].size <= 2.0{
			//	fmt.Println(fmt.Sprintf("Array [%d][%d], Coords: (%f,%f)",x,y,a,b))
				imageArr[x][y].enab='#'
				imgOut.Set(x,y, color.RGBA{0,0,0,255})
			} else {
				imageArr[x][y].enab=' '
				imgOut.Set(x,y, color.RGBA{255,255,255,255})
			}
			

//			fmt.Println(fmt.Sprintf("At end of inner y loop (%d,%d), image: %f,%f | %f @ %d",x,y,image[x][y].real,image[x][y].imag,image[x][y].size,image[x][y].count))

			//increment real component
			b += step
		}
		//increment imaginary component
		a += step
	}

	fmt.Println("calculated image, printing image...")
	var line string
	for i:=0;i<size;i++{
		line = ""
		for j := 0;j<size;j++ {
			line = line+string(imageArr[i][j].enab)
		}
		fmt.Println(line)
	}


	fmt.Println("generated png image, saving...")
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f,imgOut)

	


}
