package enigma

import (
	"log"
	"math"
)

const rotorSize = 26

type rotor struct {
	position    int
	name        string
	alphabet    string
	alphabetIdx []int
	notches     string
	notchesIdx  []int
	verbose     bool
}

func (r *rotor) step() {
	r.position = int(math.Mod(float64(r.position+26), 26)) + 1
}

func (r *rotor) cipher(input int) int {
	c := input + r.position

	c = fixAlphaIdxRange(c)
	c = r.alphabetIdx[c-1]

	if r.verbose {
		log.Printf("cipher rotor : %-5s %02d(%s) -> %02d(%s)\n",
			r.name,
			input, string(alphabetIdxToRune(alphabet, input)),
			c, string(alphabetIdxToRune(alphabet, c)))
	}
	return c
}

func (r *rotor) reverse(input int) int {
	c := input - r.position
	c = fixAlphaIdxRange(c)
	for i := 0; i < len(r.alphabetIdx); i++ {
		if r.alphabetIdx[i] == c {
			c = fixAlphaIdxRange(i + 1)
			break
		}
	}

	if r.verbose {
		log.Printf("reverse rotor: %-5s %02d(%s) -> %02d(%s)\n",
			r.name,
			input, string(alphabetIdxToRune(alphabet, input)),
			c, string(alphabetIdxToRune(alphabet, c)))
	}

	return c
}

func (r rotor) isNotch() bool {
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

func CreateRotor(name string, rotorAlphabet string, notches string, verbose bool) *rotor {
	r := rotor{
		name:        name,
		alphabet:    rotorAlphabet,
		notches:     notches,
		notchesIdx:  make([]int, len(notches)),
		alphabetIdx: make([]int, len(rotorAlphabet)),
		verbose:     verbose,
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
func RotorETW(verbose bool) rotor {
	return *CreateRotor("ETW", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "", verbose)
}

func RotorI(verbose bool) rotor {
	return *CreateRotor("I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q", verbose)
}

func RotorII(verbose bool) rotor {
	return *CreateRotor("II", "AJDKSIRUXBLHWTMCQGZNPYFVOE", "E", verbose)
}

func RotorIII(verbose bool) rotor {
	return *CreateRotor("III", "BDFHJLCPRTXVZNYEIWGAKMUSQO", "V", verbose)
}

func RotorIV(verbose bool) rotor {
	return *CreateRotor("IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB", "J", verbose)
}

func RotorV(verbose bool) rotor {
	return *CreateRotor("V", "VZBRGITYUPSDNHLXAWMJQOFECK", "Z", verbose)
}

// reflector
func RotorUKWA(verbose bool) rotor {
	return *CreateRotor("UKWA", "EJMZALYXVBWFCRQUONTSPIKHGD", "", verbose)
}

func RotorUKWB(verbose bool) rotor {
	return *CreateRotor("UKWB", "YRUHQSLDPXNGOKMIEBFZCWVJAT", "", verbose)
}

func RotorUKWC(verbose bool) rotor {
	return *CreateRotor("UKWC", "FVPJIAOYEDRZXWGCTKUQSBNMHL", "", verbose)
}
