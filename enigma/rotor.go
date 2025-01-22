package enigma

import (
	"fmt"
	"log"
	"math"
)

const rotorSize = 26

type Rotor struct {
	position     int
	ringPosition int
	name         string
	alphabet     string
	alphabetIdx  []int
	notches      string
	notchesIdx   []int
	verbose      bool
}

func (r *Rotor) setInitialPosition(initialPosition int) {
	r.position = initialPosition
}

func (r *Rotor) setRingPosition(ringPosition int) {
	r.ringPosition = ringPosition
}

func (r Rotor) String() string {
	return fmt.Sprintf("%-5s ring: %02d pos: %02d notch: %v", r.name, r.ringPosition, r.position, r.isNotch())
}

func (r *Rotor) step() {
	r.position = int(math.Mod(float64(r.position+26), 26)) + 1
}

func (r *Rotor) forward(input int) int {
	c := input - r.ringPosition + r.position

	c = fixAlphaIdxRange(c)
	alphaIdx := r.alphabetIdx[c-1]

	if r.verbose {
		log.Printf("forward %s, in: %02d(%s) out: %02d(%s)\n",
			r.String(),
			input, string(alphabetIdxToRune(alphabet, input)),
			alphaIdx, string(alphabetIdxToRune(alphabet, alphaIdx)))
	}
	return alphaIdx
}

func (r *Rotor) backward(input int) int {
	alphaIdx := input - r.ringPosition + r.position
	alphaIdx = fixAlphaIdxRange(alphaIdx)

	for i := 0; i < len(r.alphabetIdx); i++ {
		if r.alphabetIdx[i] == input {

			alphaIdx = i + 1 //- r.position
			alphaIdx = fixAlphaIdxRange(alphaIdx)
			break
		}
	}

	alphaIdx = alphaIdx + r.ringPosition - r.position
	alphaIdx = fixAlphaIdxRange(alphaIdx)
	if r.verbose {
		log.Printf("reverse %s, in: %02d(%s) out: %02d(%s)\n",
			r.String(),
			input, string(alphabetIdxToRune(alphabet, input)),
			alphaIdx, string(alphabetIdxToRune(alphabet, alphaIdx)))
	}

	return alphaIdx
}

func (r Rotor) isNotch() bool {
	for i := 0; i < len(r.notchesIdx); i++ {
		if r.notchesIdx[i] == r.position {
			return true
		}
	}
	return false
}

func fixAlphaIdxRange(idx int) int {
	if idx <= 0 {
		return idx + rotorSize
	} else if idx > rotorSize {
		return idx - rotorSize
	}
	return idx
}

//
// ----------------------------------------------------------------------------
//

func CreateRotor(name string, rotorAlphabet string, notches string, startPos int, verbose bool) *Rotor {
	r := Rotor{
		name:         name,
		alphabet:     rotorAlphabet,
		notches:      notches,
		position:     startPos - 1,
		ringPosition: 0,
		notchesIdx:   make([]int, len(notches)),
		alphabetIdx:  make([]int, len(rotorAlphabet)),
		verbose:      verbose,
	}

	if r.position < 0 {
		r.position = 0
	}

	for i, runeValue := range rotorAlphabet {
		r.alphabetIdx[i] = runeToAlphabetIdx(alphabet, runeValue)
	}

	for i, runeValue := range notches {
		r.notchesIdx[i] = runeToAlphabetIdx(alphabet, runeValue)
	}

	return &r
}

// entry wheel
func RotorETW(verbose bool) Rotor {
	return *CreateRotor("ETW", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "", 0, verbose)
}

func RotorI(verbose bool) Rotor {
	return *CreateRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q", 0, verbose)
}

func RotorII(verbose bool) Rotor {
	return *CreateRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "E", 0, verbose)
}

func RotorIII(verbose bool) Rotor {
	return *CreateRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", "V", 0, verbose)
}

func RotorIV(verbose bool) Rotor {
	return *CreateRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", 0, verbose)
}

func RotorV(verbose bool) Rotor {
	return *CreateRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", "Z", 0, verbose)
}

// reflector
func RotorUKWA(verbose bool) Rotor {
	return *CreateRotor("UKWA", "EJMZALYXVBWFCRQUONTSPIKHGD", "", 0, verbose)
}

func RotorUKWB(verbose bool) Rotor {
	return *CreateRotor("UKWB", "YRUHQSLDPXNGOKMIEBFZCWVJAT", "", 0, verbose)
}

func RotorUKWC(verbose bool) Rotor {
	return *CreateRotor("UKWC", "FVPJIAOYEDRZXWGCTKUQSBNMHL", "", 0, verbose)
}
