package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

func Pic(dx, dy int) [][]uint8 {

	ret := make([][]uint8, dy)

	for i := range ret{
		ret[i] = make([]uint8, dx)
	}

	for y, row := range ret{
		for x := range row{
			row[x] = uint8(x * y)
		}
	}



	return ret

}

func main(){
	Show(Pic)
}