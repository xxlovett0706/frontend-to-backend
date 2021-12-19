package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type gifConfig struct {
	cycles float64
	res    float64
	freq   float64
	size   int
	frames int
	delay  int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	conf := gifConfig{
		cycles: 5,
		res:    0.001,
		freq:   rand.Float64() * 3.0,
		size:   100,
		frames: 64,
		delay:  8,
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		querys := r.URL.Query()
		for i, v := range querys {
			fmt.Println(v)
			switch i {
			case "cycles":
				conf.cycles, _ = strconv.ParseFloat(v[0], 64)
			case "freq":
				conf.freq, _ = strconv.ParseFloat(v[0], 64)
			}
		}
		lissajous(rw, conf)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func lissajous(out io.Writer, conf gifConfig) {
	anim := gif.GIF{LoopCount: conf.frames}
	phase := 0.0
	for i := 0; i < conf.frames; i++ {
		rect := image.Rect(0, 0, 2*conf.size+1, 2*conf.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < conf.cycles*2*math.Pi; t += conf.res {
			x := math.Sin(t)
			y := math.Sin(t*conf.freq + phase)
			img.SetColorIndex(conf.size+int(x*float64(conf.size)+0.5), conf.size+int(y*float64(conf.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, conf.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
