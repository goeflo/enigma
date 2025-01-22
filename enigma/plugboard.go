package enigma

import (
	"log"
	"strings"
)

type Plugboard struct {
	plugs     []string
	translate map[int]int
	verbose   bool
}

func NewPlugboard(p []string, v bool) (Plugboard, error) {

	plugboard := Plugboard{
		plugs:     p,
		verbose:   v,
		translate: make(map[int]int),
	}

	for _, plug := range plugboard.plugs {
		plugboard.translate[runeToAlphabetIdx(alphabet, rune(plug[0]))] = runeToAlphabetIdx(alphabet, rune(plug[1]))
		plugboard.translate[runeToAlphabetIdx(alphabet, rune(plug[1]))] = runeToAlphabetIdx(alphabet, rune(plug[0]))
	}

	return plugboard, nil
}

func (p Plugboard) String() string {
	return strings.Join(p.plugs, ", ")

}

func (p Plugboard) forward(in int) int {

	c, ok := p.translate[in]
	if ok {
		if p.verbose {
			log.Printf("plugboard: %02d(%v) -> %02d(%v)",
				in, string(alphabetIdxToRune(alphabet, in)),
				c, string(alphabetIdxToRune(alphabet, c)))
		}
		return c
	}
	return in

}

// backward is the same because in forward and backward mappings are already in the map
func (p Plugboard) backward(in int) int {
	return p.forward(in)
}
