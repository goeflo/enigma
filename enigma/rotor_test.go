package enigma

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotor(t *testing.T) {

	r3 := RotorIII(true)
	r2 := RotorII(true)
	r1 := RotorI(true)
	r := RotorUKWB(true)

	alphaIdx := runeToAlphabetIdx(alphabet, 'G')
	cipherIdx := r3.cipher(alphaIdx) // G -> C
	cipherIdx = r2.cipher(cipherIdx) // C -> D
	cipherIdx = r1.cipher(cipherIdx) // D -> F

	cipherIdx = r.cipher(cipherIdx) // reflect F -> S

	cipherIdx = r1.reverse(cipherIdx) // S -> S
	cipherIdx = r2.reverse(cipherIdx) // S -> E
	cipherIdx = r3.reverse(cipherIdx) // E -> P

	fmt.Printf("result %v\n", cipherIdx)
	assert.Equal(t, 16, cipherIdx)

	// decrypt the cipher value
	cipherIdx = r3.cipher(cipherIdx)
	cipherIdx = r2.cipher(cipherIdx)
	cipherIdx = r1.cipher(cipherIdx)

	cipherIdx = r.cipher(cipherIdx)

	cipherIdx = r1.reverse(cipherIdx)
	cipherIdx = r2.reverse(cipherIdx)
	cipherIdx = r3.reverse(cipherIdx)

	fmt.Printf("result %v\n", cipherIdx)
	assert.Equal(t, alphaIdx, cipherIdx)

}
