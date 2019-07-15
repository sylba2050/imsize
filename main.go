package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	bounds := img.Bounds()
	fmt.Printf("width: %d, height: %d\n", bounds.Max.X, bounds.Max.Y)
	return
}
