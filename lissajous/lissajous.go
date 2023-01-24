package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var (
	verde    = color.RGBA{00, 80, 00, 5}
	vermelho = color.RGBA{255, 0, 0, 5}
	palette  = []color.Color{color.White, color.Black, verde, vermelho}
)

const (
	whiteIndex    = 0
	blackIndex    = 1
	verdeIndex    = 2
	vermelhoIndex = 3
)

func Principal() {
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Create("out2.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lissajous(f)
}

func lissajous(out io.Writer) {

	const (
		cycles  = 5     //número de revoluções completas do oscilador x
		res     = 0.001 //resolução angular
		size    = 100   //canvas da imagem cobre de [-size..+size]
		nframes = 64    //número de quadros da animação
		delay   = 8     //tempo entre quadros em unidades de 10ms
	)

	freq := rand.Float64() * 3.0 //frequência relativa do oscilador
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	corIndex := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(corIndex))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		if corIndex < 3 {
			corIndex++
		} else {
			corIndex = 0
		}
	}
	gif.EncodeAll(out, &anim)
}
