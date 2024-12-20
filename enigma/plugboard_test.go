package enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlugboard(t *testing.T) {
	p, err := NewPlugboard([]string{"AB", "CD", "EF", "GH"}, true)
	assert.Nil(t, err)

	idx := runeToAlphabetIdx(alphabet, 'A')
	crypt := p.cipher(idx)
	decrypt := p.reverse(crypt)
	assert.Equal(t, idx, decrypt)

}

func TestInvalidPlugboardConfiguration(t *testing.T) {
	_, err := NewPlugboard([]string{"AB", "CD", "AF", "GH"}, true)
	assert.NotNil(t, err)
	_, err = NewPlugboard([]string{"AB", "CH", "EF", "GH"}, true)
	assert.NotNil(t, err)
}
