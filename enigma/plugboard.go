package enigma

import (
	"log"
	"strings"
)

type Plugboard interface {
	cipher(in int) int
	String() string
}

type PlugboardImpl struct {
	plugs     []string
	translate map[int]int
	verbose   bool
}

func NewPlugboard(p []string, v bool) (Plugboard, error) {

	plugboard := PlugboardImpl{
		plugs:     p,
		verbose:   v,
		translate: make(map[int]int),
	}

	for _, plug := range plugboard.plugs {
		plugboard.translate[runeToAlphabetIdx(alphabet, rune(plug[0]))] = runeToAlphabetIdx(alphabet, rune(plug[1]))
		plugboard.translate[runeToAlphabetIdx(alphabet, rune(plug[1]))] = runeToAlphabetIdx(alphabet, rune(plug[0]))
	}

	return &plugboard, nil
}

func (p PlugboardImpl) String() string {
	return strings.Join(p.plugs, ", ")

}

func (p PlugboardImpl) cipher(in int) int {

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
	/*
		for i := 0; i < len(p.in); i++ {
			if in == p.in[i] {
				if p.verbose {
					log.Printf("plugboard: %02d(%v) -> %02d(%v)",
						in, string(alphabetIdxToRune(alphabet, in)),
						p.out[i], string(alphabetIdxToRune(alphabet, p.out[i])))
				}
				return p.out[i]
			}
		}
		return in*/
}

/*
func (p PlugboardImpl) reverse(in int) int {
	for i := 0; i < len(p.out); i++ {
		if in == p.out[i] {
			if p.verbose {
				log.Printf("plugboard: %02d(%v) -> %02d(%v)",
					in, string(alphabetIdxToRune(alphabet, in)),
					p.in[i], string(alphabetIdxToRune(alphabet, p.in[i])))
			}

			return p.in[i]
		}
	}
	return in
}
*/
