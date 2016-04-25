package glyphs

import (
	"image"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
)

// draw_one_glyph draws the named glyph onto a image canvas.
//
// The glyph name is case-insensitive.
//
// Example:
// 	DrawOneGlyph("Enlightened")
func DrawOneGlyph(glyph_name string) image.Image {
	// Initialize the graphic context on an RGBA image
	glyph_image := image.NewRGBA(image.Rect(0, 0, IMAGE_SIZE, IMAGE_SIZE))
	gc := draw2dimg.NewGraphicContext(glyph_image)

	// Set some properties
	gc.SetFillColor(color.White)
	gc.SetStrokeColor(color.Black)
	gc.SetLineCap(draw2d.RoundCap) // This doesn't work for draw2dimg. https://github.com/llgcode/draw2d/issues/27 .

	gc.SetLineWidth(HEXAGON_LINE_WIDTH)
	drawHexagon(gc)

	gc.SetLineWidth(NODE_LINE_WIDTH)
	drawNodeCircles(gc)

	gc.SetLineWidth(GLYPH_EDGE_WIDTH)
	edge_list := getEdges(glyph_name)
	drawEdges(gc, edge_list)

	drawNumber(gc, 1)

	return glyph_image
}

// draw_node_circles draws empty circles at each of the 11 nodes on the glyph pad.
func drawNodeCircles(gc draw2d.GraphicContext) {
	for _, node_char := range "0123456789a" {
		xy := nodeXy(byte(node_char), CENTRE, CENTRE, NODE_RADIUS)
		draw2dkit.Circle(gc, xy.X, xy.Y, NODE_SIZE)
		gc.FillStroke()
	}
}

// draw_hexagon draws a hexagonal border around the glyph pad.
func drawHexagon(gc draw2d.GraphicContext) {
	for n := 0; n <= 6; n++ {
		// [FIXME] 100,100,100 is hardcoded to 200px image size.
		xy := floatPolar{Mag: 1.00, Deg: float64(n)*60.0 + 30.0}.xy().moveAndScale(CENTRE, CENTRE, HEXAGON_RADIUS)
		if n == 0 {
			gc.MoveTo(xy.X, xy.Y)
		} else {
			gc.LineTo(xy.X, xy.Y)
		}
	}
	gc.FillStroke()
}

// drawEdge draws a line between nodes on the glyph pad.
//
// The line is capped with a filled circle at each end-point.
//
// node_pair must be a two-byte string containing only the letters "0123456789a", denoting the 11 possible nodes on the glyph pad.
func drawEdge(gc draw2d.GraphicContext, node_pair string) {
	xy1 := nodeXy(node_pair[0], CENTRE, CENTRE, NODE_RADIUS)
	xy2 := nodeXy(node_pair[1], CENTRE, CENTRE, NODE_RADIUS)

	gc.MoveTo(xy1.X, xy1.Y)
	gc.LineTo(xy2.X, xy2.Y)
	gc.Stroke()

	// Work around for https://github.com/llgcode/draw2d/issues/27 - drawing lines with round end-caps doesn't
	// work in draw2dimg.
	// Workaround is to manually draw little circles at each end of the line.
	gc.SetFillColor(color.Black)
	draw2dkit.Circle(gc, xy1.X, xy1.Y, NODE_SIZE) //GLYPH_EDGE_WIDTH/2)
	gc.Fill()
	draw2dkit.Circle(gc, xy2.X, xy2.Y, NODE_SIZE) //GLYPH_EDGE_WIDTH/2)
	gc.Fill()
}

// drawEdges draws all the edges specified by the edge_list argument.
// edge_list must be a string containing:
//   1. an even number of characters 0, 2, 4... and
//   2. only characters "0123456789a".
// i.e. an edge_list "2767" denoted that edges should be drawn from nodes 2-7 and nodes 6-7.
func drawEdges(gc draw2d.GraphicContext, edge_list string) {
	if len(edge_list) == 0 {
		return
	}
	if len(edge_list)%2 != 0 {
		panic("draw_edge_list() requires edge_list to have an even number of bytes.")
	}
	edge_pairs := len(edge_list) / 2
	for n := 0; n < edge_pairs; n++ {
		drawEdge(gc, edge_list[2*n:2*n+2])
	}
}