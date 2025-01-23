package enigma

import (
	"fmt"
	"testing"
)

func TestEnigmaReal(t *testing.T) {

	// U6Z DE C 1500 = 24 = WXC KCH =
	//
	// BNUGZ NIBLF MYMLL UFWCA
	// SCSSN VHAZ=

	// https://py-enigma.readthedocs.io/en/latest/overview.html
	//input := "NIBLFMYMLLUFWCASCSSNVHAZ"
	//fmt.Printf("input: %v\n", input)

	e, _ := EnigmaI([]string{"AV", "BS", "CG", "DL", "FU", "HZ", "IN", "KM", "OW", "RX"}, true)
	e.SetRotor(Right, RotorII(true))
	e.SetRotor(Middle, RotorIV(true))
	e.SetRotor(Left, RotorV(true))

	e.SetRingSettings("BUL")
	e.SetDisplay("WXC")

	fmt.Printf("%v\n", e)

	input := "KCH"
	result, _ := e.Cipher(input)
	fmt.Printf("%v -> %v\n", input, result)
	//if result != "BLA" {
	//	t.Errorf("message key should be 'BLA' but is %v", result)
	//}
	// decode BLA -> KCH

}

func TestEnigma(t *testing.T) {
	e, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
	if err != nil {
		t.Error(err)
	}

	e.SetRotor(Right, RotorI(true))
	e.SetRotor(Middle, RotorIV(true))
	e.SetRotor(Left, RotorIII(true))

	e.SetRingSettings("BUL")
	e.SetDisplay("WXC")

	fmt.Printf("%v\n", e)

	input := "HABAKUK"
	crypt, err := e.Cipher(input)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v -> %v\n", input, crypt)

	e2, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
	if err != nil {
		t.Error(err)
	}

	e2.SetRotor(Right, RotorI(true))
	e2.SetRotor(Middle, RotorIV(true))
	e2.SetRotor(Left, RotorIII(true))

	e2.SetRingSettings("BUL")
	e2.SetDisplay("WXC")

	fmt.Printf("%v\n", e2)
	crypt2, err := e2.Cipher(crypt)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v -> %v\n", crypt, crypt2)
	if input != crypt2 {
		t.Errorf("input %v and output %v should be equal", input, crypt2)
	}

}
