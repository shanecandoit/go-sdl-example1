package main

import (
	"log"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// logging
	//file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// font
	// ./assets/SourceCodePro-Regular.ttf
	var font *ttf.Font
	if err = ttf.Init(); err != nil {
		return
	}
	defer ttf.Quit()
	fontPath := "assets/SourceCodePro-Regular.ttf"
	fontSize := 16
	if font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		// return
		log.Fatal(err)
	}
	defer font.Close()
	// Create a red text with the font
	text, err := font.RenderUTF8Blended("Hello, World!", sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		return
	}
	defer text.Free()
	// Draw the text around the center of the window
	//if err = text.Blit(nil, surface, &sdl.Rect{X: 400 - (text.W / 2), Y: 300 - (text.H / 2), W: 0, H: 0}); err != nil {
	//	return
	//}

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	start := ts()
	tick := start
	frames := 0
	running := true
	var mx, my int32
	mx, my = -999, -999
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			ts := tsDiff(start)
			switch t := event.(type) {

			case *sdl.QuitEvent:

				println(ts, "Quit")
				running = false
				// break

			case *sdl.MouseMotionEvent:
				// https://github.com/veandco/go-sdl2-examples/blob/master/examples/mouse-input/mouse-input.go
				log.Println(ts, "Mouse", t.Which, "dx,dy:", t.X, t.YRel, " x,y:", t.X, t.Y)
				mx, my = t.X, t.Y
				//println("mouseMotion", x)
			}
			tick = ts
		}
		sdl.Delay(16)

		// stop after 5 seconds
		if tick > 5_000 {
			running = false
		}
		frames++

		// clear
		surface.FillRect(nil, 0)

		// draw
		rect := sdl.Rect{mx, my, 32, 32}
		surface.FillRect(&rect, 0xff00ff00)
		// draw text
		text.Blit(nil, surface, &sdl.Rect{X: 400 - (text.W / 2), Y: 300 - (text.H / 2), W: 0, H: 0})
		window.UpdateSurface()
	}
	println("drew", frames, "frames")
}

func ts() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func tsDiff(startTime int64) int64 {
	return ts() - startTime
}
