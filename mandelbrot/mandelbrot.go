package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {

	http.Handle("/mandelbrot", http.HandlerFunc(mandelbrot))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func mandelbrot(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "image/png")

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//Ponto (px, py) da imagem representa o valor complexo z
			img.Set(px, py, createMandelbrot(z))
		}
	}

	png.Encode(w, img)

}

func createMandelbrot(z complex128) color.Color {
	const interations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < interations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{249,
				214,
				133,
				contrast * n}
			//return color.Gray{255 - contrast*n}
		}
		// else {

		// }
	}

	return color.Black
}
