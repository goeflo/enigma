package enigma

import (
	"fmt"
	"testing"
)

func TestNotch(t *testing.T) {
	r1 := RotorV(true)
	r1.position = 24

	alphaIdx := runeToAlphabetIdx(alphabet, 'A')
	cipherIdx := r1.forward(alphaIdx)
	if r1.isNotch() {
		t.Errorf("there should be no notch")
	}

	r1.step()
	cipherIdx = r1.forward(cipherIdx)
	if !r1.isNotch() {
		t.Errorf("there should be a notch")
	}

	r1.step()
	cipherIdx = r1.forward(cipherIdx)
	if r1.isNotch() {
		t.Errorf("there should be no notch")
	}

	r1.step()
	_ = r1.forward(cipherIdx)
	if r1.isNotch() {
		t.Errorf("there should be no notch")
	}

}

func TestFixRange(t *testing.T) {

	for i := -5; i < 30; i++ {
		fmt.Printf(" %v %v\n", i, fixRange(i))
	}
}

func TestRotor2(t *testing.T) {

}

func TestRotor(t *testing.T) {

	ew := RotorETW(true)
	r3 := RotorIII(true)
	r2 := RotorII(true)
	r1 := RotorI(true)
	r := RotorUKWB(true)

	r3.position = 10
	r2.position = 5
	r1.position = 17

	// crypt
	alphaIdx := runeToAlphabetIdx(alphabet, 'G')

	cipherIdx := ew.forward(alphaIdx)

	cipherIdx = r3.forward(cipherIdx) // G -> C
	cipherIdx = r2.forward(cipherIdx) // C -> D
	cipherIdx = r1.forward(cipherIdx) // D -> F

	cipherIdx = r.forward(cipherIdx) // reflect F -> S

	cipherIdx = r1.backward(cipherIdx) // S -> S
	cipherIdx = r2.backward(cipherIdx) // S -> E
	cipherIdx = r3.backward(cipherIdx) // E -> P

	cipherIdx = ew.backward(cipherIdx)

	fmt.Printf("result %v\n", cipherIdx)
	//assert.Equal(t, 16, cipherIdx)

	// decrypt
	cipherIdx = ew.forward(cipherIdx)
	cipherIdx = r3.forward(cipherIdx)
	cipherIdx = r2.forward(cipherIdx)
	cipherIdx = r1.forward(cipherIdx)

	cipherIdx = r.forward(cipherIdx)

	cipherIdx = r1.backward(cipherIdx)
	cipherIdx = r2.backward(cipherIdx)
	cipherIdx = r3.backward(cipherIdx)
	cipherIdx = ew.backward(cipherIdx)

	fmt.Printf("result %v\n", cipherIdx)
	if alphaIdx != cipherIdx {
		t.Errorf("input char %v does not match result char %v", alphaIdx, cipherIdx)
	}

}

func TestReverse(t *testing.T) {
	r := RotorIII(true)
	r.position = 4
	alphaIdx := runeToAlphabetIdx(alphabet, 'B')
	cipherIdx := r.forward(alphaIdx)
	decrypt := r.backward(cipherIdx)

	if runeToAlphabetIdx(alphabet, 'B') != decrypt {
		t.Errorf("start idx %v does not match result idx %v", runeToAlphabetIdx(alphabet, 'B'), decrypt)
	}

}

func TestWithRing(t *testing.T) {
	r := RotorIII(true)
	r.position = 4
	r.ringPosition = 20

	alphaIdx := runeToAlphabetIdx(alphabet, 'B')
	cipherIdx := r.forward(alphaIdx)
	decrypt := r.backward(cipherIdx)

	if runeToAlphabetIdx(alphabet, 'B') != decrypt {
		t.Errorf("start idx %v does not match result idx %v", runeToAlphabetIdx(alphabet, 'B'), decrypt)
	}

	alphaIdx = runeToAlphabetIdx(alphabet, 'Q')
	cipherIdx = r.forward(alphaIdx)
	decrypt = r.backward(cipherIdx)

	if runeToAlphabetIdx(alphabet, 'Q') != decrypt {
		t.Errorf("start idx %v does not match result idx %v", runeToAlphabetIdx(alphabet, 'Q'), decrypt)
	}
}

func TestNotched(t *testing.T) {

	r := RotorII(true)
	r.step()
	r.step()
	if r.isNotch() {
		t.Errorf("there should be no notch")
	}

	r.step()
	if r.isNotch() {
		t.Errorf("there should be no notch")
	}

	r.step()
	if !r.isNotch() {
		t.Errorf("there should be a notch")
	}

	r.step()
	if r.isNotch() {
		t.Errorf("there should be no notch")
	}

}

/*
{2, 'A', 'C', false, false, 17, 24}, // Q -> X
		{1, 'C', 'D', false, false, 21, 19}, // U -> S
		{1, 'D', 'E', false, true, 5, 20},   // E -> T
		{1, 'E', 'F', true, false, 5, 23},   // E -> W
		{1, 'F', 'G', false, false, 14, 8},  // N -> H
*/
