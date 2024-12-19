package enigma

import (
	"fmt"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Enigma struct {
	version    string
	entryRotor rotor
	rotors     []rotor
	reflector  rotor
	plugboard  Plugboard
	verbose    bool
}

func EnigmaI(v bool) Enigma {
	e := Enigma{
		version:    "EnigmaI",
		entryRotor: RotorETW(v),
		reflector:  RotorUKWB(v),
		rotors:     make([]rotor, 3),
		verbose:    v,
	}

	return e
}

func (e Enigma) String() string {
	return fmt.Sprintf("enigma: %v, entry rotor: %v, reflector: %v", e.version, e.entryRotor.name, e.reflector.name)
}

func (e *Enigma) AddRotor(p int, r rotor) error {
	if p < 1 || p > len(e.rotors) {
		return fmt.Errorf("rotor index out of range, 1 <= %v <= %v", p, len(e.rotors))
	}
	e.rotors[p-1] = r
	return nil
}

func (e *Enigma) Decrypt(input string) (string, error) {
	return "", fmt.Errorf("not yet implemented")
}

func alphabetIdxToRune(a string, i int) rune {
	return rune(a[i-1])
}

func runeToAlphabetIdx(a string, r rune) int {
	for i, runeValue := range a {
		if r == runeValue {
			return i + 1
		}
	}
	return -1
}
