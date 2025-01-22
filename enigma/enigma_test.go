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

	result, _ := e.Cipher("KCH")
	fmt.Printf("result: %v\n", result)
	if result != "BLA" {
		t.Errorf("message key should be 'BLA' but is %v", result)
	}
	// decode BLA -> KCH

}

//func TestEnigma(t *testing.T) {
//	e, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
//	assert.Nil(t, err)
//
//	err = e.SetRotor(Right, RotorI(1, true))
//	assert.Nil(t, err)
//
//	err = e.SetRotor(Middle, RotorIV(1, true))
//	assert.Nil(t, err)
//
//	err = e.SetRotor(Left, RotorIII(1, true))
//	assert.Nil(t, err)
//
//	fmt.Printf("%v\n", e)
//
//	input := "HABAKUK"
//	crypt, err := e.Cipher(input)
//	assert.Nil(t, err)
//	fmt.Printf("%v -> %v\n", input, crypt)
//	fmt.Printf("%v\n", e)
//
//	e2, err := EnigmaI([]string{"HZ", "YR", "IF", "QT", "JN", "GC", "AP", "UX", "BD", "KS"}, true)
//	assert.Nil(t, err)
//
//	err = e2.SetRotor(Right, RotorI(1, true))
//	assert.Nil(t, err)
//
//	err = e2.SetRotor(Middle, RotorIV(1, true))
//	assert.Nil(t, err)
//
//	err = e2.SetRotor(Left, RotorIII(1, true))
//	assert.Nil(t, err)
//
//	fmt.Printf("%v\n", e2)
//	crypt2, err := e2.Cipher(crypt)
//	assert.Nil(t, err)
//	fmt.Printf("%v -> %v\n", crypt, crypt2)
//	assert.Equal(t, input, crypt2)
//}
//
