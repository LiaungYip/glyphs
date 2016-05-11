package glyphs

import (
	"strings"
	"testing"
)

var spellingTests = []struct {
	in  string
	out string
}{
	{"xm", "XM"}, // Correct spellings should return themselves.
	{"Xm", "XM"},
	{"xM", "XM"},
	{"XM", "XM"},
	{"Portals", "Portal"}, // Plural / not plural
	{"Humans", "Human"},
	{"Shaper", "Shapers"},
	{"Protal", "Portal"},             // Transposition
	{"Porttal", "Portal"},            // Double character
	{"Prtal", "Portal"},              // Missing character
	{"wipqozn", ""},                  // Should fail.
	{"xm xm", ""},                    // Missing comma
	{"Civilisation", "Civilization"}, // American spelling

	// Actual errors from user testing
	{"/glyph success", ""}, // /glyphs success would have worked
	{"doorway", ""},
	{"enlightened have mind journey", ""},
	{"@IngressGlyphBot unbounded", ""},
	{"Civilisation", "Civilization"},
	{"understand", ""},
	{"human failure", ""},
	{"bathtub", ""},
	{"left kite", ""},
	{"right ear", ""},
	{"shapers destroy humans", ""},
}

func TestSpellcheck(t *testing.T) {
	for _, tt := range spellingTests {
		out := Spellcheck(tt.in)
		if strings.Compare(tt.out, out) != 0 {
			t.Errorf("Input `%s` -> `%s`, expected `%s`", tt.in, out, tt.out)
		}
	}
}
