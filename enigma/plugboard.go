package enigma

type Plugboard struct {
}

func Crypt(in rune) int {
	// TODO
	return runeToAlphabetIdx(alphabet, in)
}
