package glyphs

import (
	"testing"
	"github.com/llgcode/draw2d/draw2dimg"

	"fmt"
	"sort"
	"strings"
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

var lookupTests = []struct {
	in  string
	glyphName string
	edgeList string
}{
	{"XM","XM","67697a898a"}, // Test that correct capitalisation is returned
	{"xm","XM","67697a898a"},
	{"Xm","XM","67697a898a"},
	{"xM","XM","67697a898a"},
	{"OPEN ALL","Open All","010512233437384578"}, // Test that things with spaces are accurately returned
	{"enlightened","Enlightened","01091223696a9a"},
	{"n'zeer","N'zeer","06090a3a6a9a"},
	{"clear all","Clear All","01050a1223343a45"}, // Regression test. The "Clear All" alias for "Clearall" wasn't present in original data.
	// Sustain
	// Sustain All
}

func TestLookup(t *testing.T) {
	for _, tt := range lookupTests{
		g, e, err := Lookup(tt.in)
		if err != nil {
			t.Errorf("%s -> error! %s", tt.in, err.Error())
			continue
		}
		if (g != tt.glyphName) || (e != tt.edgeList) {
			t.Errorf("%s -> %s, %s. Wanted %s, %s.", tt.in, g, e, tt.glyphName, tt.edgeList)
		}
		t.Logf("Case `%s` -> `%s`, `%s` passed!", tt.in, tt.glyphName, tt.edgeList)
	}
}

// TODO: Test that all glyph edge lists in the dictionary are in 'canonical form'

func TestCanonicalOrdering(t *testing.T) {
	for edgeList, glyphNames := range Glyphs {
		var pairs []string
		for i := 0; i < len(edgeList)/2; i++ {
			pair := edgeList[i*2:i*2+2]
			if pair[0] > pair[1] {
				pair = string(pair[1]) + string(pair[0])
			}
			pairs = append(pairs, pair)
		}
		sort.Strings(pairs)
		newlist := strings.Join(pairs,"")
		if strings.Compare(edgeList, newlist) != 0 {
			fmt.Println("Not canonical!", edgeList, "should be", newlist, glyphNames)
		}
		fmt.Println(edgeList)
	}
}