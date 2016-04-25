package glyphs

const RT3 = 1.731

const IMAGE_SIZE = 512.0
const CENTRE = IMAGE_SIZE / 2
const NODE_RADIUS = IMAGE_SIZE / 2 * 0.75
const NODE_SIZE = IMAGE_SIZE * 0.04
const HEXAGON_RADIUS = IMAGE_SIZE / 2 * 0.95

const NODE_LINE_WIDTH = IMAGE_SIZE * 0.01
const HEXAGON_LINE_WIDTH = IMAGE_SIZE * 0.01
const GLYPH_EDGE_WIDTH = IMAGE_SIZE * 0.03

type settings struct {
	imageSize        float64
	centre           float64
	nodeRadius       float64
	nodeSize         float64
	hexagonRadius    float64
	nodeLineWidth    float64
	hexagonLineWidth float64
	glyphEdgeWidth   float64
}

func defaultSettings (imageSize float64) settings {
	var s settings

	s.imageSize = imageSize

	s.centre = s.imageSize / 2
	s.nodeRadius = s.imageSize / 2 * 0.75
	s.nodeSize = s.imageSize * 0.04
	s.hexagonRadius = s.imageSize / 2 * 0.95

	s.nodeLineWidth = s.imageSize * 0.01
	s.hexagonLineWidth = s.imageSize * 0.01
	s.glyphEdgeWidth = s.imageSize * 0.03

	return s
}