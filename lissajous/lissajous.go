package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

//go build .\main.go

func main() {
	Principal()
}

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

type ValuesLissajous struct {
	Cycles  int     //número de revoluções completas do oscilador x
	Res     float64 //resolução angular
	Size    int     //canvas da imagem cobre de [-size..+size]
	NFrames int     //número de quadros da animação
	Delay   int     //tempo entre quadros em unidades de 10ms
}

func Principal() {
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Create("arquivos/out2.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	Lissajous(f, ValuesLissajous{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		NFrames: 64,
		Delay:   8,
	})
}

func Lissajous(out io.Writer, values ValuesLissajous) {

	freq := rand.Float64() * 3.0 //frequência relativa do oscilador
	anim := gif.GIF{LoopCount: values.NFrames}
	phase := 0.0
	corIndex := 0
	for i := 0; i < values.NFrames; i++ {
		rect := image.Rect(0, 0, 2*values.Size+1, 2*values.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(values.Cycles)*2*math.Pi; t += values.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(values.Size+int(x*float64(values.Size)+0.5), values.Size+int(y*float64(values.Size)+0.5), uint8(corIndex))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, values.Delay)
		anim.Image = append(anim.Image, img)

		if corIndex < 3 {
			corIndex++
		} else {
			corIndex = 0
		}
	}
	gif.EncodeAll(out, &anim)
}
