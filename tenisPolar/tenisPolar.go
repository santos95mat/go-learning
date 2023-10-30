package tenisPolar

func Crypto(str string) string {
	tp := []string{
		"t", "e", "n", "i", "s", "p", "o", "l", "a", "r",
		"T", "E", "N", "I", "S", "P", "O", "L", "A", "R",
		"é", "è", "ê", "ë", "ó", "ò", "ô", "õ",
		"í", "ì", "î", "ï", "á", "à", "â", "ã",
	}

	pt := []string{
		"p", "o", "l", "a", "r", "t", "e", "n", "i", "s",
		"P", "O", "L", "A", "R", "T", "E", "N", "I", "S",
		"ó", "ò", "ô", "õ", "é", "è", "ê", "ë",
		"á", "à", "â", "ã", "í", "ì", "î", "ï",
	}

	crypto := ""

	for j, l := range str {
		for i, m := range tp {
			if string(l) == m {
				crypto += pt[i]
			}
		}

		if len(crypto) == j {
			crypto += string(l)
		}
	}

	return crypto
}
