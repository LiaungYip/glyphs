package glyphs

// getEdges looks up a glyph name in the glyph dictionary, and returns the edge_list used to draw that glyph.
// The glyph_name is case-insensitive, but must match the spelling exactly.
// 2016-04-28: This function is deprecated. Superceded by Lookup().
func getEdges(glyph_name string) string {
	_, edge_list, err := Lookup(glyph_name)
	if err != nil {
		return ""
	} else {
		return edge_list
	}
}