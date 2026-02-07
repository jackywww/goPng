package main

import "github.com/fogleman/gg"

func main() {
	saveImage("C", 20, "Icon-App-20x20@1x.png", 22.0, false, "image/")
	saveImage("C", 40, "Icon-App-20x20@2x.png", 44.0, false, "image/")
	saveImage("C", 60, "Icon-App-20x20@3x.png", 60.0, false, "image/")
	saveImage("C", 29, "Icon-App-29x29@1x.png", 30.0, false, "image/")
	saveImage("C", 58, "Icon-App-29x29@2x.png", 58.0, false, "image/")
	saveImage("C", 87, "Icon-App-29x29@3x.png", 86.0, false, "image/")
	saveImage("C", 40, "Icon-App-40x40@1x.png", 44.0, false, "image/")
	saveImage("C", 80, "Icon-App-40x40@2x.png", 84.0, false, "image/")
	saveImage("C", 120, "Icon-App-40x40@3x.png", 113.0, false, "image/")
	saveImage("C", 120, "Icon-App-60x60@2x.png", 113.0, false, "image/")
	saveImage("C", 180, "Icon-App-60x60@3x.png", 165.0, false, "image/")
	saveImage("C", 76, "Icon-App-76x76@1x.png", 81.0, false, "image/")
	saveImage("C", 152, "Icon-App-76x76@2x.png", 152.0, false, "image/")
	saveImage("C", 167, "Icon-App-83.5x83.5@2x.png", 156.0, false, "image/")
	saveImage("C", 1024, "Icon-App-1024x1024@1x.png", 970.0, false, "image/")

}

func saveImage(letter string, imageSize int, name string, fontSize float64, channel bool, path string) {
	S := imageSize
	dc := gg.NewContext(S, S)
	if channel {
		dc.SetRGBA(0, 0, 0, 0)
	} else {
		dc.SetRGB(1, 1, 1)
	}

	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("gilmer-heavy.ttf", fontSize); err != nil {
		panic(err)
	}
	textWidth, textHeight := dc.MeasureString(letter)

	x := (float64(imageSize) - textWidth) / 2
	y := ((float64(imageSize)-textHeight)/2 + textHeight) * 0.99

	dc.DrawString(letter, x, y)
	dc.SavePNG(path + name)
}
