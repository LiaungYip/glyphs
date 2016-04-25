package glyphs

import (
	"testing"
	"github.com/llgcode/draw2d/draw2dimg"

)

func TestDrawOneGlyph(t *testing.T) {
	s := DefaultSettings(100)
	glyph_image := DrawOneGlyph("Enlightened", s)
	draw2dimg.SaveToPngFile("./sample/enlightened.png", glyph_image)
}

func TestDrawGlyphSequence(t *testing.T) {
	s := DefaultSettings(100)
	glyph_names := []string{"help", "enlightened", "capture", "all", "portal"}
	glyph_image := DrawGlyphSequence(glyph_names, s)
	draw2dimg.SaveToPngFile("./sample/help_enlightened_capture_all_portal.png", glyph_image)
}