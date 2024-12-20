package enigma

import (
	"fmt"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Enigma interface {
	Cipher(input string) (string, error)
	SetRotor(p int, r rotor, rotorPos int) error
}

type EnigmaImpl struct {
	version    string
	entryRotor rotor
	rotors     []rotor
	reflector  rotor
	plugboard  Plugboard
	verbose    bool
}

func EnigmaI(plugs []string, v bool) (Enigma, error) {

	p, err := NewPlugboard(plugs, v)
	if err != nil {
		return nil, err
	}

	e := &EnigmaImpl{
		version:    "EnigmaI",
		entryRotor: RotorETW(1, v),
		reflector:  RotorUKWB(1, v),
		rotors:     make([]rotor, 3),
		plugboard:  p,
		verbose:    v,
	}

	return e, nil
}

func (e *EnigmaImpl) String() string {
	return fmt.Sprintf("enigma: %v, entry rotor: %v, reflector: %v\nplugboard: %v\nrotor configuration: %v(%02d) %v(%02d) %v(%02d)\n",
		e.version, e.entryRotor.name, e.reflector.name, e.plugboard.String(),
		e.rotors[0].name, e.rotors[0].position, e.rotors[1].name, e.rotors[1].position, e.rotors[2].name, e.rotors[2].position)
}

// p is the position of the rotor in the enigma, for EnigmaI 1-3
// rotorPosition is the current rotor position
func (e *EnigmaImpl) SetRotor(p int, r rotor, rotorPosition int) error {
	if p < 1 || p > len(e.rotors) {
		return fmt.Errorf("rotor index out of range, 1 <= %v <= %v", p, len(e.rotors))
	}
	for _, er := range e.rotors {
		if r.name == er.name {
			return fmt.Errorf("rotor %v already in the enigma", r.name)
		}
	}
	e.rotors[p-1] = r
	return nil
}

func (e EnigmaImpl) Cipher(in string) (string, error) {
	in = strings.ToUpper(in)
	crypt := ""
	for _, runeVal := range in {
		out := e.plugboard.cipher(runeToAlphabetIdx(alphabet, runeVal))

		out = e.entryRotor.cipher(out)

		for i := 0; i < len(e.rotors); i++ {
			if i == 0 {
				e.rotors[i].step()
			} else {
				if e.rotors[i-1].isNotch() {
					e.rotors[i].step()
				}
			}
			out = e.rotors[i].cipher(out)
		}

		out = e.reflector.cipher(out)

		for i := len(e.rotors) - 1; i >= 0; i-- {
			out = e.rotors[i].reverse(out)
		}

		out = e.entryRotor.reverse(out)

		out = e.plugboard.reverse(out)

		crypt += string(alphabetIdxToRune(alphabet, out))
	}

	return crypt, nil
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
