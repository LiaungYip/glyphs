package glyphs

// All settings are in pixels.
type Settings struct {
	ImageSize        int
	Centre           float64
	CodeRadius       float64
	NodeSize         float64
	HexagonRadius    float64
	NodeLineWidth    float64
	HexagonLineWidth float64
	GlyphEdgeWidth   float64
}

// DefaultSettings gives sane, good-looking defaults for the size and line width
// of the glyph drawings.
func DefaultSettings (imageSize int) Settings {
	var s Settings

	s.ImageSize = imageSize

	is := float64(imageSize)

	s.Centre = is / 2.0
	s.CodeRadius = is / 2 * 0.75
	s.NodeSize = is * 0.04
	s.HexagonRadius = is / 2 * 0.95

	s.NodeLineWidth = is * 0.01
	s.HexagonLineWidth = is * 0.01
	s.GlyphEdgeWidth = is * 0.03

	return s
}