package glyphs

import (
	"github.com/llgcode/draw2d"
	"image"
	"image/draw"
	"math"
)



// draw_glyph_sequence draws the named glyphs onto an image canvas.
//
// If there are more than two glyphs, the glyphs are shown on multiple
// rows and columns.
func DrawGlyphSequence(glyph_names []string) image.Image {
	var images []image.Image

	for _, glyph_name := range glyph_names {
		images = append(images, DrawOneGlyph(glyph_name))
	}
	return squarePackImages(images)
}

// squarePackImages packs multiple glyph images onto one image.
func squarePackImages(images []image.Image) image.Image {
	num_glyphs := len(images)

	// Work out how many rows and columns to display the glyphs in.
	// Try and keep it square. If the number of rows and columns aren't even, prefer "wide" instead of "tall".
	num_glyphs_float := float64(num_glyphs)
	array_width := math.Ceil(math.Sqrt(num_glyphs_float))
	array_height := math.Ceil(num_glyphs_float / array_width)

	rows := int(array_height)
	cols := int(array_width)

	SZ := int(IMAGE_SIZE)
	combined_image := image.NewRGBA(image.Rect(0, 0, SZ*cols, SZ*rows))

	for n := 0; n < num_glyphs; n++ {
		col := n % cols
		row := n / cols // Integer division rounds down
		glyph_image := images[n]
		dest_rect := image.Rect(0, 0, SZ, SZ).Add(image.Pt(col*SZ, row*SZ))
		draw.Draw(combined_image, dest_rect, glyph_image, image.Pt(0, 0), draw.Src)
	}
	return combined_image
}

func drawNumber(gc draw2d.GraphicContext, number int) {
	gc.SetFontData(draw2d.FontData{Name: "Times New Roman"})
	gc.FillStringAt("HELLO", 10, 10)
}