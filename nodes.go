package glyphs

// Node Indices:
//         0
//   5            1
//      9      6
//         10
//      8      7
//   4            2
//         3
//
var nodes = [11]floatPolar{
	{Mag: 1.0, Deg: 90.0},   // 0
	{Mag: 1.0, Deg: 30.0},   // 1
	{Mag: 1.0, Deg: -30.0},  // 2
	{Mag: 1.0, Deg: -90.0},  // 3
	{Mag: 1.0, Deg: -150.0}, // 4
	{Mag: 1.0, Deg: 150.0},  // 5
	{Mag: 0.5, Deg: 30.0},   // 6
	{Mag: 0.5, Deg: -30.0},  // 7
	{Mag: 0.5, Deg: -150.0}, // 8
	{Mag: 0.5, Deg: 150.0},  // 9
	{Mag: 0.0, Deg: 0.0},    // 10 / "a"
}

// nodeXy returns the pixel coordinates of a node on the glyph pad.
// This includes accounting for translation and scaling.
//
// i.e. nodeXy(0) returns the coordinates of the top-most node.
func nodeXy(node_char byte, origin_x, origin_y, scale float64) xy {
	node_number, ok := nodeChars[node_char]
	if !ok {
		panic("node_char not found in NODE_CHARS mapping. It should be in the set '0123456789a'.")
	}
	return nodes[node_number].xy().moveAndScale(origin_x, origin_y, scale)
}