package enigma

import (
	"fmt"
	"testing"
)

func TestPlugboard(t *testing.T) {
	p, err := NewPlugboard([]string{"AB", "CD", "EF", "GH"}, true)
	if err != nil {
		t.Error(err)
	}

	idx := runeToAlphabetIdx(alphabet, 'A')
	crypt := p.forward(idx)
	fmt.Printf("%v\n", crypt)
	decrypt := p.forward(crypt)
	if idx != decrypt {
		t.Errorf("input %v and result %v should be the same", idx, decrypt)
	}

}
