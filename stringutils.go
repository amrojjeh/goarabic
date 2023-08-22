// Package goarabic contains utility functions for working with Arabic strings.
package goarabic

import (
	"errors"
	"fmt"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// SmartLength returns the length of the given string
// without considering the Arabic Vowels (Tashkeel).
func SmartLength(s *string) int {
	// len() use int as return value, so we'd better follow for compatibility
	length := 0

	for _, value := range *s {
		if tashkeel[value] {
			continue
		}
		length++
	}

	return length
}

// RemoveTashkeel returns its argument as rune-wise string without Arabic vowels (Tashkeel).
func RemoveTashkeel(s string) string {
	// var r []rune
	// the capcity of the slice wont be greater than the length of the string itself
	// hence, cap = len(s)
	r := make([]rune, 0, len(s))

	for _, value := range s {
		if tashkeel[value] {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

// RemoveTatweel returns its argument as rune-wise string without Arabic Tatweel character.
func RemoveTatweel(s string) string {
	r := make([]rune, 0, len(s))

	for _, value := range s {
		if TATWEEL.equals(value) {
			continue
		}
		r = append(r, value)
	}

	return string(r)
}

func getCharGlyph(previousChar, currentChar, nextChar rune) rune {
	glyph := currentChar
	previousIn := false // in the Arabic Alphabet or not
	nextIn := false     // in the Arabic Alphabet or not

	for _, s := range alphabet {
		if s.equals(previousChar) { // previousChar in the Arabic Alphabet ?
			previousIn = true
		}

		if s.equals(nextChar) { // nextChar in the Arabic Alphabet ?
			nextIn = true
		}
	}

	for _, s := range alphabet {

		if !s.equals(currentChar) { // currentChar in the Arabic Alphabet ?
			continue
		}

		if previousIn && nextIn { // between two Arabic Alphabet, return the medium glyph
			for s := range beggining_after {
				if s.equals(previousChar) {
					return getHarf(currentChar).Beginning
				}
			}

			return getHarf(currentChar).Medial
		}

		if nextIn { // beginning (because the previous is not in the Arabic Alphabet)
			return getHarf(currentChar).Beginning
		}

		if previousIn { // final (because the next is not in the Arabic Alphabet)
			for s := range beggining_after {
				if s.equals(previousChar) {
					return getHarf(currentChar).Isolated
				}
			}
			return getHarf(currentChar).Final
		}

		if !previousIn && !nextIn {
			return getHarf(currentChar).Isolated
		}

	}
	return glyph
}

// equals() return if true if the given Arabic char is alphabetically equal to
// the current Harf regardless its shape (Glyph)
func (c *Harf) equals(char rune) bool {
	switch char {
	case c.Base:
		return true
	case c.Beginning:
		return true
	case c.Isolated:
		return true
	case c.Medial:
		return true
	case c.Final:
		return true
	default:
		return false
	}
}

// getHarf gets the correspondent Harf for the given Arabic char
func getHarf(char rune) Harf {
	for _, s := range alphabet {
		if s.equals(char) {
			return s
		}
	}

	return Harf{Base: char, Isolated: char, Medial: char, Final: char}
}

//RemoveAllNonAlphabetChars deletes all characters which are not included in Arabic Alphabet
func RemoveAllNonArabicChars(text string) string {
	runes := []rune(text)
	newText := []rune{}
	for _, current := range runes {
		inAlphabet := false
		for _, s := range alphabet {
			if s.equals(current) {
				inAlphabet = true
			}
		}
		if inAlphabet {
			newText = append(newText, current)
		}
	}
	return string(newText)
}

// ToGlyph returns the glyph representation of the given text
func ToGlyph(text string) string {
	var prev, next rune

	runes := []rune(text)
	length := len(runes)
	newText := make([]rune, 0, length)

	for i, current := range runes {
		// get the previous char
		if (i - 1) < 0 {
			prev = 0
		} else {
			prev = runes[i-1]
		}

		// get the next char
		if (i + 1) <= length-1 {
			next = runes[i+1]
		} else {
			next = 0
		}

		// get the current char representation or return the same if unnecessary
		glyph := getCharGlyph(prev, current, next)

		// append the new char representation to the newText
		newText = append(newText, glyph)
	}

	return string(newText)
}

// ArToSafeBWLetter converts Arabic text to the Safe Buckwalter encoding scheme,
// which can be found here: https://camel-tools.readthedocs.io/en/stable/reference/encoding_schemes.html
func ArToSafeBWLetter(r rune) (rune, error) {
	switch r {
	case HAMZA.Base:
		return 'C', nil
	case ALEF_MADDA_ABOVE.Base:
		return 'M', nil
	case ALEF_HAMZA_ABOVE.Base:
		return 'O', nil
	case WAW_HAMZA_ABOVE.Base:
		return 'W', nil
	case ALEF_HAMZA_BELOW.Base:
		return 'I', nil
	case YEH_HAMZA_ABOVE.Base:
		return 'Q', nil
	case ALEF.Base:
		return 'A', nil
	case BEH.Base:
		return 'b', nil
	case TEH_MARBUTA.Base:
		return 'p', nil
	case TEH.Base:
		return 't', nil
	case THEH.Base:
		return 'v', nil
	case JEEM.Base:
		return 'j', nil
	case HAH.Base:
		return 'H', nil
	case KHAH.Base:
		return 'x', nil
	case DAL.Base:
		return 'd', nil
	case THAL.Base:
		return 'V', nil
	case REH.Base:
		return 'r', nil
	case SEEN.Base:
		return 's', nil
	case SHEEN.Base:
		return 'c', nil
	case SAD.Base:
		return 'S', nil
	case DAD.Base:
		return 'D', nil
	case TAH.Base:
		return 'T', nil
	case ZAH.Base:
		return 'Z', nil
	case AIN.Base:
		return 'E', nil
	case GHAIN.Base:
		return 'g', nil
	case TATWEEL.Base:
		return '_', nil
	case FEH.Base:
		return 'f', nil
	case QAF.Base:
		return 'q', nil
	case KAF.Base:
		return 'k', nil
	case LAM.Base:
		return 'l', nil
	case MEEM.Base:
		return 'm', nil
	case NOON.Base:
		return 'n', nil
	case HEH.Base:
		return 'h', nil
	case WAW.Base:
		return 'w', nil
	case ALEF_MAKSURA.Base:
		return 'Y', nil
	case YEH.Base:
		return 'y', nil
	case FATHATAN:
		return 'F', nil
	case DAMMATAN:
		return 'N', nil
	case KASRATAN:
		return 'K', nil
	case FATHA:
		return 'a', nil
	case DAMMA:
		return 'u', nil
	case KASRA:
		return 'i', nil
	case SHADDA:
		return '~', nil
	case SUKUN:
		return 'o', nil
	case SUPERSCRIPT_ALEF:
		return 'e', nil
	case PEH.Base:
		return 'P', nil
	case TCHEH.Base:
		return 'J', nil
	case VEH.Base:
		return 'B', nil
	case GAF.Base:
		return 'G', nil
	case ' ':
		return ' ', nil
	default:
		return 0, errors.New(fmt.Sprintf("%c is not a valid Arabic character", r))
	}
}

func ArToSafeBW(s string) (string, error) {
	result := ""
	for _, v := range s {
		bw, err := ArToSafeBWLetter(v)
		if err != nil {
			return "", err
		}
		result += string(bw)
	}
	return result, nil
}

func SafeBWToArLetter(r rune) (rune, error) {
	switch r {
	case 'C':
		return HAMZA.Base, nil
	case 'M':
		return ALEF_MADDA_ABOVE.Base, nil
	case 'O':
		return ALEF_HAMZA_ABOVE.Base, nil
	case 'W':
		return WAW_HAMZA_ABOVE.Base, nil
	case 'I':
		return ALEF_HAMZA_BELOW.Base, nil
	case 'Q':
		return YEH_HAMZA_ABOVE.Base, nil
	case 'A':
		return ALEF.Base, nil
	case 'b':
		return BEH.Base, nil
	case 'p':
		return TEH_MARBUTA.Base, nil
	case 't':
		return TEH.Base, nil
	case 'v':
		return THEH.Base, nil
	case 'j':
		return JEEM.Base, nil
	case 'H':
		return HAH.Base, nil
	case 'x':
		return KHAH.Base, nil
	case 'd':
		return DAL.Base, nil
	case 'V':
		return THAL.Base, nil
	case 'r':
		return REH.Base, nil
	case 's':
		return SEEN.Base, nil
	case 'c':
		return SHEEN.Base, nil
	case 'S':
		return SAD.Base, nil
	case 'D':
		return DAD.Base, nil
	case 'T':
		return TAH.Base, nil
	case 'Z':
		return ZAH.Base, nil
	case 'E':
		return AIN.Base, nil
	case 'g':
		return GHAIN.Base, nil
	case '_':
		return TATWEEL.Base, nil
	case 'f':
		return FEH.Base, nil
	case 'q':
		return QAF.Base, nil
	case 'k':
		return KAF.Base, nil
	case 'l':
		return LAM.Base, nil
	case 'm':
		return MEEM.Base, nil
	case 'n':
		return NOON.Base, nil
	case 'h':
		return HEH.Base, nil
	case 'w':
		return WAW.Base, nil
	case 'Y':
		return ALEF_MAKSURA.Base, nil
	case 'y':
		return YEH.Base, nil
	case 'F':
		return FATHATAN, nil
	case 'N':
		return DAMMATAN, nil
	case 'K':
		return KASRATAN, nil
	case 'a':
		return FATHA, nil
	case 'u':
		return DAMMA, nil
	case 'i':
		return KASRA, nil
	case '~':
		return SHADDA, nil
	case 'o':
		return SUKUN, nil
	case 'e':
		return SUPERSCRIPT_ALEF, nil
	case 'P':
		return PEH.Base, nil
	case 'J':
		return TCHEH.Base, nil
	case 'B':
		return VEH.Base, nil
	case 'G':
		return GAF.Base, nil
	case ' ':
		return ' ', nil
	default:
		return 0, errors.New(fmt.Sprintf("%c is not a valid Safe BW character", r))
	}
}

func SafeBWToAr(s string) (string, error) {
	result := ""
	for _, v := range s {
		ar, err := SafeBWToArLetter(v)
		if err != nil {
			return "", err
		}
		result += string(ar)
	}
	return result, nil
}
