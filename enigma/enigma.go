package enigma

import (
	"fmt"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

/* TODO do we need this interface
type Part interface {
	forward(in int) int
	backward(in int) int
}*/

type RotorPosition int

const (
	Right RotorPosition = iota
	Middle
	Left
)

type Enigma interface {
	Cipher(input string) (string, error)
	SetRotor(p RotorPosition, r Rotor) error
	SetRingSettings(ringSettings string) error
	SetDisplay(startPos string) error
}

type EnigmaImpl struct {
	version    string
	entryRotor Rotor
	rotors     []Rotor
	reflector  Rotor
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
		entryRotor: RotorETW(v),
		reflector:  RotorUKWB(v),
		rotors:     make([]Rotor, 3),
		plugboard:  p,
		verbose:    v,
	}

	return e, nil
}

func (e *EnigmaImpl) String() string {
	return fmt.Sprintf("enigma: %v, entry rotor: %v, reflector: %v\nplugboard: %v\nrotor configuration: %v(%02d/%c) %v(%02d/%c) %v(%02d/%c)\nring position      : %v(%02d/%c) %v(%02d/%c) %v(%02d/%c)\n",
		e.version, e.entryRotor.name, e.reflector.name, e.plugboard.String(),
		e.rotors[0].name, e.rotors[0].position, alphabetIdxToRune(e.rotors[0].alphabet, e.rotors[0].position),
		e.rotors[1].name, e.rotors[1].position, alphabetIdxToRune(e.rotors[1].alphabet, e.rotors[1].position),
		e.rotors[2].name, e.rotors[2].position, alphabetIdxToRune(e.rotors[2].alphabet, e.rotors[2].position),
		e.rotors[0].name, e.rotors[0].ringPosition, alphabetIdxToRune(e.rotors[0].alphabet, e.rotors[0].ringPosition),
		e.rotors[1].name, e.rotors[1].ringPosition, alphabetIdxToRune(e.rotors[1].alphabet, e.rotors[1].ringPosition),
		e.rotors[2].name, e.rotors[2].ringPosition, alphabetIdxToRune(e.rotors[2].alphabet, e.rotors[2].ringPosition))
}

func (e *EnigmaImpl) SetRingSettings(ringSettings string) error {

	if len(ringSettings) < len(e.rotors) {
		return fmt.Errorf("not enough ring settings, we have %v rotos but only %v ring settings", len(e.rotors), len(ringSettings))
	}

	for k, rune := range ringSettings {
		e.rotors[k].setInitialPosition(runeToAlphabetIdx(e.rotors[k].alphabet, rune))
	}
	return nil
}

func (e *EnigmaImpl) SetDisplay(startPos string) error {

	if len(startPos) < len(e.rotors) {
		return fmt.Errorf("not enough ring settings, we have %v rotos but only %v ring settings", len(e.rotors), len(startPos))
	}

	for k, rune := range startPos {
		e.rotors[k].setRingPosition(runeToAlphabetIdx(e.rotors[k].alphabet, rune))
	}
	return nil
}

// p is the position of the rotor in the enigma, for EnigmaI 1-3
// rotorPosition is the current rotor position
func (e *EnigmaImpl) SetRotor(p RotorPosition, r Rotor) error {
	/*if p < 1 || p > len(e.rotors) {
		return fmt.Errorf("rotor index out of range, 1 <= %v <= %v", p, len(e.rotors))
	}*/
	for _, er := range e.rotors {
		if r.name == er.name {
			return fmt.Errorf("rotor %v already in the enigma", r.name)
		}
	}
	e.rotors[p] = r
	return nil
}

func (e EnigmaImpl) Cipher(in string) (string, error) {
	in = strings.ToUpper(in)
	crypt := ""
	for _, runeVal := range in {

		for i := 1; i < len(e.rotors); i++ {
			if e.rotors[i-1].isNotch() {
				e.rotors[i].step()
			}
		}
		e.rotors[0].step()

		//		out := e.entryRotor.forward(runeToAlphabetIdx(alphabet, runeVal))
		//		out = e.rotors[0].forward(out)
		out := e.plugboard.forward(runeToAlphabetIdx(alphabet, runeVal))
		out = e.entryRotor.forward(out)

		for i := 0; i < len(e.rotors); i++ {
			out = e.rotors[i].forward(out)
		}

		out = e.reflector.forward(out)

		for i := len(e.rotors) - 1; i >= 0; i-- {
			out = e.rotors[i].backward(out)
		}

		out = e.entryRotor.backward(out)
		out = e.plugboard.forward(out)

		crypt += string(alphabetIdxToRune(alphabet, out))
	}

	return crypt, nil
}

func alphabetIdxToRune(a string, i int) rune {
	return rune(a[i])
}

func runeToAlphabetIdx(a string, r rune) int {
	for i, runeValue := range a {
		if r == runeValue {
			return i
		}
	}
	return -1
}
