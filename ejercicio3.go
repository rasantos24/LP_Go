package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"golang.org/x/image/bmp"
)

func main() {
	bitmap, fe := openImage("Kaiba.bmp")
	if fe != nil {
		fmt.Println(fe)
	}

	limi := bitmap.Bounds()
	ancho, altura := limi.Max.X, limi.Max.Y

	myBMP := image.NewRGBA(limi)
	for y := 0; y < altura; y++ {
		for x := 0; x < ancho; x++ {
			pu := bitmap.At(x, y)
			red, green, blue, _ := pu.RGBA()
			rgb := (red + green + blue) / 3
			pi := color.Gray{uint8(rgb / 256)}
			myBMP.Set(x, y, pi)
		}
	}

	archiOUT, fe := os.Create("Kaiba_2.bmp")
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
