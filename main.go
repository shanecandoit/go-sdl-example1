package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
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
	for running {
		mx, my = -999, -999
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			ts := tsDiff(start)
			switch t := event.(type) {

			case *sdl.QuitEvent:

				println(ts, "Quit")
				running = false
				// break

			case *sdl.MouseMotionEvent:
				// https://github.com/veandco/go-sdl2-examples/blob/master/examples/mouse-input/mouse-input.go
				fmt.Println(ts, "Mouse", t.Which, "dx,dy:", t.X, t.YRel, " x,y:", t.X, t.Y)
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
