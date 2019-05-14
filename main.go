package main

/*
	@author	Boosh
	@desc	Program to generate some pixel map of the mandelbrot set
*/

import (
	"fmt"
	"math"
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
	printTextOutput := false

	//edge size of square for mandelbrot image, and mandelbrot image
	//the array sizes, and square length must be size according to how the algo was written
	size := 2000
	var imageArr [2000][2000] complex
	imgOut := image.NewNRGBA(image.Rect(0,0,2000,2000))

	//step of 4 corresonds to a -2 to +2 range on the real and imaginary components
	step := 4.0/float64(size)

	//fun values of c
	c := complex{real:-1.2,imag:0.1}
	//c := complex{real:-0.95,imag:0.07}

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

			//"enab" can be removed as it's for the initial text-mapping but doesn't work with higher resolutions
			//can also make this logical bit more granular for different colors and artsiness
			//	i.e. size of 2 but different count values, different values under 2, etc.
			if imageArr[x][y].size <= 2.0{
				imageArr[x][y].enab='#'
				imgOut.Set(x,y, color.RGBA{0,0,0,255})
			} else {
				imageArr[x][y].enab=' '
				imgOut.Set(x,y, color.RGBA{255,255,255,255})
			}
			//increment real component
			b += step
		}
		//increment imaginary component
		a += step
	}

	if printTextOutput {
		fmt.Println("calculated image, printing image...")
		var line string
		for i:=0;i<size;i++{
			line = ""
			for j := 0;j<size;j++ {
				line = line+string(imageArr[i][j].enab)
			}
			fmt.Println(line)
		}
	}

	fmt.Println("generated png image, saving...")
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f,imgOut)
}
