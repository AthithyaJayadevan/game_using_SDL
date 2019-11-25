package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const width = 800
const height = 600

type color struct {
	r, g, b byte
}

func setpixel(x, y int, c color, pixel []byte) {
	index := (y*width + x) + 4

	if index < len(pixel)-4 && index >= 0 {
		pixel[index] = c.r
		pixel[index+1] = c.g
		pixel[index+2] = c.b
	}

}

func main() {

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("An error happend :%v", err)
		os.Exit(1)
	}
	defer window.Destroy()
	renderfunc, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Something happend , error: %v", err)
		os.Exit(1)
	}
	defer renderfunc.Destroy()

	texture, err := renderfunc.CreateTexture(sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING, int32(width), int32(height))

	if err != nil {
		fmt.Printf("Something happend , error: %v", err)
		os.Exit(1)
	}
	defer texture.Destroy()

	pixels := make([]byte, width*height*4)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			setpixel(j, i, color{255, 0, 0}, pixels)
		}
	}

	texture.Update(nil, pixels, width*4)
	renderfunc.Copy(texture, nil, nil)
	renderfunc.Present()
	sdl.Delay(5000)

}
