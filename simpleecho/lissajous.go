package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{
	color.White,
	color.RGBA{231, 20, 73, 1},
	color.RGBA{0, 102, 0, 1},
	color.RGBA{38, 20, 231, 1},
}

// Lissajous draw a figure
func Lissajous(out io.Writer) {
	const (
		cycles  = 5     // кол-во полных колебаний
		res     = 0.001 // угловое разрешение
		size    = 100   // канва изображения охватывает [size..+size]
		nframes = 64    // количество кадров анимации
		delay   = 8     // задержка между кадрами (1 == 10мс)
	)

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		nextIndex := uint8(rand.Intn(4-1) + 1) // случайный выбор цвета из палитры (1..3)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size*0.5), size+int(y*size+0.5), nextIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // игнорируем ошибки
}
