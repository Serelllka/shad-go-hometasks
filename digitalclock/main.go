package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/now", handler)
	port := flag.Int("port", 8080, "a port")
	flag.Parse()
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(*port), nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	k, _ := strconv.Atoi(req.URL.Query().Get("k"))
	if k < 0 || k > 30 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	k = int(math.Max(float64(k), 1))
	stringTime := req.URL.Query().Get("time")
	if stringTime == "" {
		stringTime = time.Now().Format("15:04:05")
	}
	tm, err := time.Parse("15:04:05", stringTime)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	img := genImageFromTime(tm, k)
	err = png.Encode(w, img)

	if err != nil {
		return
	}
}

func genImageFromTime(tm time.Time, k int) (img *image.RGBA) {
	height := (len([]rune(Zero)) + 1) / (strings.Index(Zero, "\n") + 1)
	width := strings.Index(Zero, "\n")*6 + strings.Index(Colon, "\n")*2
	img = image.NewRGBA(image.Rect(0, 0, width*k, height*k))
	stringTime := []rune(tm.Format("15:04:05"))
	offset := 0
	for _, lit := range stringTime {
		currLit, _ := GetCurrentLiteralPerf(lit)
		currWidth := strings.Index(string(currLit), "\n")
		for i := 0; i < height; i++ {
			for j := 0; j < currWidth; j++ {
				if currLit[i*(currWidth+1)+j] == '1' {
					FillCord(offset+j, i, k, img, Cyan)
				}
			}
		}
		offset += currWidth
	}
	return img
}

func GetCurrentLiteralPerf(literal rune) (rn []rune, err error) {
	err = nil
	rn = []rune{}
	switch literal {
	case '0':
		rn = []rune(Zero)
	case '1':
		rn = []rune(One)
	case '2':
		rn = []rune(Two)
	case '3':
		rn = []rune(Three)
	case '4':
		rn = []rune(Four)
	case '5':
		rn = []rune(Five)
	case '6':
		rn = []rune(Six)
	case '7':
		rn = []rune(Seven)
	case '8':
		rn = []rune(Eight)
	case '9':
		rn = []rune(Nine)
	case ':':
		rn = []rune(Colon)
	}
	return rn, err
}

func FillCord(x, y, k int, img *image.RGBA, clr color.Color) {
	realX := x * k
	realY := y * k
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			img.Set(realX+i, realY+j, clr)
		}
	}
}
