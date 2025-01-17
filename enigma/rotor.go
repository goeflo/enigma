package enigma

import (
	"fmt"
	"log"
	"math"
)

const rotorSize = 26

type rotor struct {
	offset      int
	name        string
	alphabet    string
	alphabetIdx []int
	notches     string
	notchesIdx  []int
	verbose     bool
}

func (r rotor) String() string {
	return fmt.Sprintf("%-5s pos: %02d", r.name, r.offset)
}

func (r *rotor) step() {
	r.offset = int(math.Mod(float64(r.offset+26), 26)) + 1
}

func (r *rotor) cipher(input int) int {
	c := input + r.offset

	c = fixAlphaIdxRange(c)
	c = r.alphabetIdx[c-1]

	if r.verbose {
		log.Printf("cipher  %s, %02d(%s) -> %02d(%s)\n",
			r.String(),
			input, string(alphabetIdxToRune(alphabet, input)),
			c, string(alphabetIdxToRune(alphabet, c)))
	}
	return c
}

func (r *rotor) reverse(input int) int {
	//c := input - r.position

	//c = fixAlphaIdxRange(c)
	//fmt.Printf("after pos %v\n", c)
	c := input

	for i := 0; i < len(r.alphabetIdx); i++ {
		if r.alphabetIdx[i] == input {

			c = i + 1 - r.offset
			c = fixAlphaIdxRange(c)
			//fmt.Printf("index %v -pos %v alph %v\n", i+1, i+1-r.position, r.alphabetIdx[i-r.position])
			//c := input - r.position
			//c = fixAlphaIdxRange(c)
			break
		}
	}

	if r.verbose {
		log.Printf("reverse %s, %02d(%s) -> %02d(%s)\n",
			r.String(),
			input, string(alphabetIdxToRune(alphabet, input)),
			c, string(alphabetIdxToRune(alphabet, c)))
	}

	return c
}

func (r rotor) isNotch() bool {
	for i := 0; i < len(r.notchesIdx); i++ {
		if r.notchesIdx[i] == r.offset {
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

func CreateRotor(name string, rotorAlphabet string, notches string, startPos int, verbose bool) *rotor {
	r := rotor{
		name:        name,
		alphabet:    rotorAlphabet,
		notches:     notches,
		offset:      startPos - 1,
		notchesIdx:  make([]int, len(notches)),
		alphabetIdx: make([]int, len(rotorAlphabet)),
		verbose:     verbose,
	}

	if r.offset < 0 {
		r.offset = 0
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
func RotorETW(startPos int, verbose bool) rotor {
	return *CreateRotor("ETW", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "", startPos, verbose)
}

func RotorI(startPos int, verbose bool) rotor {
	return *CreateRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q", startPos, verbose)
}

func RotorII(startPos int, verbose bool) rotor {
	return *CreateRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "E", startPos, verbose)
}

func RotorIII(startPos int, verbose bool) rotor {
	return *CreateRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", "V", startPos, verbose)
}

func RotorIV(startPos int, verbose bool) rotor {
	return *CreateRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", startPos, verbose)
}

func RotorV(startPos int, verbose bool) rotor {
	return *CreateRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", "Z", startPos, verbose)
}

// reflector
func RotorUKWA(startPos int, verbose bool) rotor {
	return *CreateRotor("UKWA", "EJMZALYXVBWFCRQUONTSPIKHGD", "", startPos, verbose)
}

func RotorUKWB(startPos int, verbose bool) rotor {
	return *CreateRotor("UKWB", "YRUHQSLDPXNGOKMIEBFZCWVJAT", "", startPos, verbose)
}

func RotorUKWC(startPos int, verbose bool) rotor {
	return *CreateRotor("UKWC", "FVPJIAOYEDRZXWGCTKUQSBNMHL", "", startPos, verbose)
}
