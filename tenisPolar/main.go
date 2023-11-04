package tenisPolar

var tp = [52]string{
	"t", "e", "n", "i", "s", "p", "o", "l", "a", "r",
	"T", "E", "N", "I", "S", "P", "O", "L", "A", "R",
	"é", "è", "ê", "ë", "ó", "ò", "ô", "õ",
	"É", "È", "Ê", "Ë", "Ó", "Ò", "Ô", "Õ",
	"í", "ì", "î", "ï", "á", "à", "â", "ã",
	"Í", "Ì", "Î", "Ï", "Á", "À", "Â", "Ã",
}

var pt = [52]string{
	"p", "o", "l", "a", "r", "t", "e", "n", "i", "s",
	"P", "O", "L", "A", "R", "T", "E", "N", "I", "S",
	"ó", "ò", "ô", "õ", "é", "è", "ê", "ë",
	"Ó", "Ò", "Ô", "Õ", "É", "È", "Ê", "Ë",
	"á", "à", "â", "ã", "í", "ì", "î", "ï",
	"Á", "À", "Â", "Ã", "Í", "Ì", "Î", "Ï",
}

func Crypto(str string) string {
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
