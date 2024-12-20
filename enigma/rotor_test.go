package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotch(t *testing.T) {
	r1 := RotorV(1, true)
	r1.position = 24

	alphaIdx := runeToAlphabetIdx(alphabet, 'A')
	cipherIdx := r1.cipher(alphaIdx)
	r1.step()
	assert.False(t, r1.isNotch())

	cipherIdx = r1.cipher(cipherIdx)
	r1.step()
	assert.True(t, r1.isNotch())

	_ = r1.cipher(cipherIdx)
	r1.step()
	assert.False(t, r1.isNotch())

}

func TestRotor(t *testing.T) {

	ew := RotorETW(1, true)
	r3 := RotorIII(1, true)
	r2 := RotorII(10, true)
	r1 := RotorI(1, true)
	r := RotorUKWB(1, true)

	// crypt
	alphaIdx := runeToAlphabetIdx(alphabet, 'G')

	cipherIdx := ew.cipher(alphaIdx)

	cipherIdx = r3.cipher(cipherIdx) // G -> C
	cipherIdx = r2.cipher(cipherIdx) // C -> D
	cipherIdx = r1.cipher(cipherIdx) // D -> F

	cipherIdx = r.cipher(cipherIdx) // reflect F -> S

	cipherIdx = r1.reverse(cipherIdx) // S -> S
	cipherIdx = r2.reverse(cipherIdx) // S -> E
	cipherIdx = r3.reverse(cipherIdx) // E -> P

	cipherIdx = ew.reverse(cipherIdx)

	fmt.Printf("result %v\n", cipherIdx)
	//assert.Equal(t, 16, cipherIdx)

	// decrypt
	cipherIdx = ew.cipher(cipherIdx)
	cipherIdx = r3.cipher(cipherIdx)
	cipherIdx = r2.cipher(cipherIdx)
	cipherIdx = r1.cipher(cipherIdx)

	cipherIdx = r.cipher(cipherIdx)

	cipherIdx = r1.reverse(cipherIdx)
	cipherIdx = r2.reverse(cipherIdx)
	cipherIdx = r3.reverse(cipherIdx)
	cipherIdx = ew.reverse(cipherIdx)

	fmt.Printf("result %v\n", cipherIdx)
	assert.Equal(t, alphaIdx, cipherIdx)

}

func TestReverse(t *testing.T) {
	r := RotorIII(5, true)
	alphaIdx := runeToAlphabetIdx(alphabet, 'B')
	cipherIdx := r.cipher(alphaIdx)
	decrypt := r.reverse(cipherIdx)

	assert.Equal(t, runeToAlphabetIdx(alphabet, 'B'), decrypt)

}
