package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
	blue  color.Color = color.RGBA{0, 0, 255, 255}
)

// ref) http://golang.org/doc/articles/image_draw.html
func main() {

	m := image.NewRGBA(image.Rect(0, 0, 640, 480)) //*NRGBA (image.Image interface)

	// fill m in blue
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	// draw a line
	for i := m.Bounds().Min.X; i < m.Bounds().Max.X; i++ {
		m.Set(i, m.Bounds().Max.Y/2, white) // to change a single pixel
	}

	w, _ := os.Create("new.png")
	defer w.Close()
	png.Encode(w, m) //Encode writes the Image m to w in PNG format.

	Show(w.Name())

}

// show  a specified file by Preview.app for OS X(darwin)
func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
