package main

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

func main() {
	bitMap, fe := openImage("Kaiba.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	limi := bitMap.Bounds()
	myBMP := image.NewRGBA(image.Rect(0, 0, 128, 128))

	ancho, altura := limi.Max.X/128, limi.Max.X/128

	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			pixel := bitMap.At(x*ancho, y*altura)
			myBMP.Set(x, y, pixel)
		}
	}

	archiOUT, fe := os.Create("Mini_Kaiba.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	defer archiOUT.Close()

	bmp.Encode(archiOUT, myBMP)

}

func openImage(rafaam string) (image.Image, error) {
	file, fe := os.Open(rafaam)
	if fe != nil {
		return nil, fe
	}
	defer file.Close()
	return bmp.Decode(file)
}
