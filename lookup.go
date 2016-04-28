package glyphs

import (
	"strings"
	"errors"
	"fmt"
)

// Lookup find a glyph's canonical name and edge list, from some user-inputted name.
// Returns (glyph's canonical name, glyph's edge list, error code).
func Lookup(glyph_name string) (string, string, error) {
	for edge_list, glyph_names := range Glyphs {
		for _, g := range glyph_names {
			// EqualFold: test if strings are equal under Unicode case-folding.
			if strings.EqualFold(glyph_name, g) {
				return g, edge_list, nil
			}
		}
	}
	return "","",errors.New(fmt.Sprintf("Glyph named `%s` not found in dictionary.", glyph_name))
}
