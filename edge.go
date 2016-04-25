package glyphs

import (
	"strings"
)



// getEdges looks up a glyph name in the glyph dictionary, and returns the edge_list used to draw that glyph.
// The glyph_name is case-insensitive, but must match the spelling exactly.
// Todo: add Levenshtein Distance so minor typographical errors can be corrected.
func getEdges(glyph_name string) string {
	UP := strings.ToUpper
	glyph_name = UP(glyph_name)
	for edge_list, glyph_names := range Glyphs {
		for _, g := range glyph_names {
			if glyph_name == UP(g) {
				return edge_list
			}
		}
	}
	return ""
}