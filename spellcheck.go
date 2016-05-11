package glyphs

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"sort"
	"strings"
	"log"
)

// Define a `corrections` slice type, satisfying sort.Interface.

type correction struct {
	w        string
	distance int
}

type corrections []correction

func (slice corrections) Len() int {
	return len(slice)
}

func (slice corrections) Less(i, j int) bool {
	return slice[i].distance < slice[j].distance
}

func (slice corrections) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Spellcheck compares the spelling of the input word, to the words in the glyph dictionary.

func Spellcheck(in string) (string) {
	// Check if input word is in our dictionary.
	for _, g := range GlyphNames {
		if strings.EqualFold(in, g) {
			return g
		}
	}
	// Input word isn't in our dictionary.

	var candidates corrections
	options := levenshtein.DefaultOptions
	for _, g := range GlyphNames {
		g_lower := []rune(strings.ToLower(g))
		in_lower := []rune(strings.ToLower(in))
		distance := levenshtein.DistanceForStrings(g_lower, in_lower, options)
		if distance < 3 { // Edit distance of 3 allows for one or two typos.
			candidates = append(candidates, correction{g, distance})
			log.Printf("Candidate correction: `%s` -> `%s` (distance: %d)\n", in, g, distance)
		}
	}
	if candidates.Len() > 0 {
		sort.Sort(candidates)
		return candidates[0].w
	} else {
		log.Printf("No candidate corrections for `%s`", in)
		return ""
	}
}
