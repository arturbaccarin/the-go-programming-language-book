package exercise

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
	"net/url"
	"strconv"
)

/*
Exercise 1.12 - Modify the Lissajous server to read parameter values from the URL.
For example, you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
number of cycles to 20 instead of the default 5.
*/
func Exercise1_12() {
	http.HandleFunc("/lissajous", handlerLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Lissajous(cycles float64, out io.Writer) {
	var palette = []color.Color{color.White, color.Black}

	const (
		res        = 0.001 // angular resolution
		size       = 100   // image canvas covers [-size..+size]
		nframes    = 64    // number of animation frames
		delay      = 8     // delay between frames in 10ms units
		whiteIndex = 0     // first color in palette
		blackIndex = 1     // next color in palette
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func handlerLissajous(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, "Failed to parse URL", http.StatusBadRequest)
		return
	}

	cyclesParam := u.Query().Get("cycles")

	if cyclesParam != "" {
		cyclesFloat, err := strconv.ParseFloat(cyclesParam, 64)
		if err != nil {
			http.Error(w, "Failed to convert 'cyclesParam' to float", http.StatusBadRequest)
			return
		}

		Lissajous(cyclesFloat, w)

	} else {
		fmt.Fprint(w, "Cycles parameter not provided")
	}
}
