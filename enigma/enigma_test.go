package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnigma(t *testing.T) {
	e, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
	assert.Nil(t, err)

	err = e.SetRotor(1, RotorI(1, true), 1)
	assert.Nil(t, err)

	err = e.SetRotor(2, RotorIV(1, true), 1)
	assert.Nil(t, err)

	err = e.SetRotor(3, RotorIII(1, true), 1)
	assert.Nil(t, err)

	fmt.Printf("%v\n", e)

	input := "HABAKUK"
	crypt, err := e.Cipher(input)
	assert.Nil(t, err)
	fmt.Printf("%v -> %v\n", input, crypt)
	fmt.Printf("%v\n", e)

	e2, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
	assert.Nil(t, err)

	err = e2.SetRotor(1, RotorI(1, true), 1)
	assert.Nil(t, err)

	err = e2.SetRotor(2, RotorIV(1, true), 1)
	assert.Nil(t, err)

	err = e2.SetRotor(3, RotorIII(1, true), 1)
	assert.Nil(t, err)

	fmt.Printf("%v\n", e2)
	crypt2, err := e2.Cipher(crypt)
	assert.Nil(t, err)
	fmt.Printf("%v -> %v\n", crypt, crypt2)
}
