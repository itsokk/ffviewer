package main

import (
	"bytes"
	"encoding/binary"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MAGIC = "farbfeld"
)

type FarbfeldHeader struct {
	magic  [8]byte
	width  uint32
	height uint32
}

type FarbfeldPixel struct {
	R uint16
	G uint16
	B uint16
	A uint16
}

type FarbfeldImage struct {
	header FarbfeldHeader
	data   []FarbfeldPixel
}

func farbfeldDecoder(data []byte) FarbfeldImage {
	if !bytes.Equal(data[:8], []byte(MAGIC)) {
		panic("Not a farbfeld image.")
	}
	width := binary.BigEndian.Uint32(data[8:12])
	height := binary.BigEndian.Uint32(data[12:16])
	image := FarbfeldImage{
		FarbfeldHeader{},
		make([]FarbfeldPixel, width*height),
	}

	copy(image.header.magic[:], MAGIC)
	image.header.width = width
	image.header.height = height

	for i := 0; i < int(width*height)*8; i += 8 {
		image.data[i/8] = FarbfeldPixel{
			binary.BigEndian.Uint16(data[16+i : 16+i+2]),
			binary.BigEndian.Uint16(data[16+i+2 : 16+i+4]),
			binary.BigEndian.Uint16(data[16+i+4 : 16+i+6]),
			binary.BigEndian.Uint16(data[16+i+6 : 16+i+8]),
		}
	}

	return image
}

func main() {
	if len(os.Args) < 2 {
		panic("No filename specified.")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	image := farbfeldDecoder(data)

	rl.InitWindow(int32(image.header.width), int32(image.header.height), os.Args[1])
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for y := uint32(0); y < image.header.height; y++ {
			for x := uint32(0); x < image.header.width; x++ {
				index := y*image.header.width + x
				color := rl.Color{
					R: uint8(image.data[index].R >> 8),
					G: uint8(image.data[index].G >> 8),
					B: uint8(image.data[index].B >> 8),
					A: uint8(image.data[index].A >> 8),
				}

				rl.DrawRectangle(int32(x), int32(y), 1, 1, color)
			}
		}

		rl.EndDrawing()
	}
}
