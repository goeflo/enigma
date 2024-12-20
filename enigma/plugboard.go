package enigma

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/maps"
)

type Plugboard interface {
	cipher(in int) int
	reverse(in int) int
	String() string
}

type PlugboardImpl struct {
	plugs   []string
	in      []int
	out     []int
	verbose bool
}

func NewPlugboard(p []string, v bool) (Plugboard, error) {

	plugboard := PlugboardImpl{
		plugs:   p,
		verbose: v,
		in:      make([]int, len(p)),
		out:     make([]int, len(p)),
	}

	for i, plug := range plugboard.plugs {
		plugboard.in[i] = runeToAlphabetIdx(alphabet, rune(plug[0]))
		plugboard.out[i] = runeToAlphabetIdx(alphabet, rune(plug[1]))
	}

	if !plugboard.isPlugConfigValid() {
		return nil, fmt.Errorf("invalid plugboard configuration")
	}

	return &plugboard, nil
}

func (p PlugboardImpl) String() string {
	strb := strings.Builder{}
	for i := 0; i < len(p.plugs); i++ {
		strb.WriteString(fmt.Sprintf("%s(%02d,%02d)", p.plugs[i], p.in[i], p.out[i]))
	}
	return strb.String()
}

func (p PlugboardImpl) isPlugConfigValid() bool {

	tin := map[int]int{}
	tout := map[int]int{}

	for i := 0; i < len(p.in); i++ {
		tin[p.in[i]] = 0
		tout[p.out[i]] = 0
	}

	if len(p.in) != len(maps.Keys(tin)) || (len(p.in) != len(maps.Keys(tout))) {
		return false
	}
	return true

}

func (p PlugboardImpl) cipher(in int) int {
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
	return in
}

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
