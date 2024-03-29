// Lissajous generates GIF animations of random Lissajous figures.
package main

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
	bgColor = color.Black
	red     = color.RGBA{0xff, 0, 0, 0xff}
	orange  = color.RGBA{0xff, 0xf5, 0, 0xff}
	yellow  = color.RGBA{0xff, 0xff, 0, 0xff}
	green   = color.RGBA{0, 0xff, 0, 0xff}
	aqua    = color.RGBA{0, 0xff, 0xff, 0xff}
	blue    = color.RGBA{0, 0, 0xff, 0xff}
	purple  = color.RGBA{0x80, 0, 0x80, 0xff}
	palette = []color.Color{bgColor, red, orange, yellow, green, aqua, blue, purple}
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions 何周するか
		res    = 0.001 // angular resolution 角度分解能
		size   = 100   // image canvas covers [-size..+size]
		nframe = 64    // number of animation frames
		delay  = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator 相対周波数
	anim := gif.GIF{LoopCount: nframe}
	numsOfColors := len(palette[1:]) // 背景色以外の色の数
	phase := 0.0                     // phase difference 位相差
	for i := 0; i < nframe; i++ {
		// フレーム毎に色を変化させる
		lineIndex := i%numsOfColors + 1

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(lineIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding error
}
