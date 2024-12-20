package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlugboard(t *testing.T) {
	p, err := NewPlugboard([]string{"AB", "CD", "EF", "GH"}, true)
	assert.Nil(t, err)

	idx := runeToAlphabetIdx(alphabet, 'A')
	crypt := p.cipher(idx)
	fmt.Printf("%v\n", crypt)
	decrypt := p.cipher(crypt)
	assert.Equal(t, idx, decrypt)

}
