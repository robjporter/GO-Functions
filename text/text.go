package text

import (
	"bytes"
	"errors"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/robjporter/go-functions/as"
)

const UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const VOWELS = "aeoui"
const CONSONANTS = "bcdfghjklmnpqrstvwxyz"
const NUMBERS = "1234567890"
const SPECIALS = "!@#$%^&*-_"
const MAXPREFIXLENGTH = 4

var boostThreshold float64 = 0.7
var prefixScale float64 = 0.1
var noop = func(a rune) rune { return a }

var military = map[string]string{
	" ": " | ",
	"a": "alfa",
	"b": "bravo",
	"c": "charlie",
	"d": "delta",
	"e": "echo",
	"f": "foxtrot",
	"g": "golf",
	"h": "hotel",
	"i": "india",
	"j": "juliet",
	"k": "kilo",
	"l": "lima",
	"m": "mike",
	"n": "november",
	"o": "oscar",
	"p": "papa",
	"q": "quebec",
	"r": "romeo",
	"s": "sierra",
	"t": "tango",
	"u": "uniform",
	"v": "victor",
	"w": "whiskey",
	"x": "x-ray",
	"y": "yankee",
	"z": "zulu",
}

var morse = map[string]string{
	"a":  ". _",
	"b":  "_ . . .",
	"c":  "_ . _ .",
	"d":  "_ . .",
	"e":  ".",
	"f":  ". . _ .",
	"g":  "_ _ .",
	"h":  ". . . .",
	"i":  ". .",
	"j":  ". _ _ _",
	"k":  "_ . _",
	"l":  ". _ . .",
	"m":  "_ _",
	"n":  "_ .",
	"o":  "_ _ _",
	"p":  ". _ _ .",
	"q":  "_ _ . _",
	"r":  ". _ .",
	"s":  ". . .",
	"t":  "_",
	"u":  ". . _",
	"v":  ". . . _",
	"w":  ". _ _",
	"x":  "_ . . _",
	"y":  "_ . _ _",
	"z":  "_ _ . .",
	"0":  "_ _ _ _ _",
	"1":  ". _ _ _ _",
	"2":  ". . _ _ _",
	"3":  ". . . _ _",
	"4":  ". . . . _",
	"5":  ". . . . .",
	"6":  "_ . . . .",
	"7":  "_ _ . . .",
	"8":  "_ _ _ . .",
	"9":  "_ _ _ _ .",
	".":  "·–·–·– ",
	",":  "––··–– ",
	"?":  "··––·· ",
	"'":  "·––––· ",
	"!":  "–·–·–– ",
	"/":  "–··–· ",
	"(":  "–·––· ",
	")":  "–·––·– ",
	"&":  "·–··· ",
	":":  "–––··· ",
	";":  "–·–·–· ",
	"=":  "–···– ",
	"+":  "·–·–· ",
	"–":  "–····– ",
	"_":  "··––·– ",
	"\"": "·–··–· ",
	"$":  "···–··– ",
	"@":  "·––·–· ",
}

//Morse////////////////////////////////////////////////////////////////////
func ToMorse(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		ret += morse[strings.ToLower(s[i:i+1])] + " | "
	}
	return ret
}

//Military/////////////////////////////////////////////////////////////////
func ToMilitary(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		ret += military[strings.ToLower(s[i:i+1])] + " | "
	}
	return ret
}

//File/////////////////////////////////////////////////////////////////////
func FileExtension(str string) string {
	for i := len(str) - 1; i > -1; i-- {
		if str[i] == '.' {
			return str[i+1 : len(str)]
		}
	}
	return ""
}

//Title////////////////////////////////////////////////////////////////////
func ToTitle(str string) string {
	arrWords := strings.Split(str, " ")
	newString := func(str string, r rune, i int) string {
		return str[:i] + string(r) + str[i+1:]
	}
	var buffer bytes.Buffer
	for i := 0; i < len(arrWords); i++ {
		if arrWords[i][0] >= 'a' && arrWords[i][0] <= 'z' {
			buffer.WriteString(newString(arrWords[i], rune(arrWords[i][0])-32, 0))
			buffer.WriteString(" ")
		} else {
			buffer.WriteString(arrWords[i])
			buffer.WriteString(" ")
		}
	}
	return buffer.String()
}

//Train////////////////////////////////////////////////////////////////////
func ToTrain(s string) string {
	return snaker(s, '-', unicode.ToUpper, unicode.ToUpper, noop)
}

//Spinal////////////////////////////////////////////////////////////////////
func ToSpinal(s string) string {
	return snaker(s, '-', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

//Camel to Snake////////////////////////////////////////////////////////////
func ToSnake(s string) string {
	return snaker(s, '_', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

func ToSnakeUpper(s string) string {
	return snaker(s, '_', unicode.ToUpper, unicode.ToUpper, unicode.ToUpper)
}

//Snake to Camel////////////////////////////////////////////////////////////
func ToCamel(s string) string {
	return snaker(s, rune(0), unicode.ToUpper, unicode.ToUpper, noop)
}

func ToCamelLower(s string) string {
	return snaker(s, rune(0), unicode.ToLower, unicode.ToUpper, noop)
}

//Reverse//////////////////////////////////////////////////////////////////
func ToReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//Random String////////////////////////////////////////////////////////////
func GenerateRandomString(length int) string {
	return generateRandom(length, false)
}

func GenerateRandomStringSpecial(length int) string {
	return generateRandom(length, true)
}

func generateRandomContains(str string, num int) bool {
	l := strings.ContainsAny(LOWERCASE, str)
	u := strings.ContainsAny(UPPERCASE, str)
	n := strings.ContainsAny(NUMBERS, str)
	s := strings.ContainsAny(SPECIALS, str)

	if num == 3 {
		return l && u && n
	} else if num == 4 {
		return l && u && n && s
	}
	return false
}

func generateRandom(length int, specials bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := make([]string, length)
	choiceset := ""
	if specials {
		choiceset = LOWERCASE + UPPERCASE + NUMBERS + SPECIALS
	} else {
		choiceset = LOWERCASE + UPPERCASE + NUMBERS
	}

	for i := 0; i < length; i++ {
		index := r.Intn(len(choiceset))
		str[i] = choiceset[index : index+1]
	}

	tmpStr := strings.Join(str, "")

	num := 3
	if specials {
		num = 4
	}

	if generateRandomContains(tmpStr, num) {
		return tmpStr
	} else {
		generateRandom(length, specials)
	}
	return ""
}

//jaro///////////////////////////////////////////////////////////////////////
func Jaro(s1, s2 string) float64 {
	m1 := commonCharacters(s1, s2)
	m2 := commonCharacters(s2, s1)
	lm1 := float64(len(m1))
	lm2 := float64(len(m2))
	// divide transpositions by 2 (as required by the algorithm)
	t := float64(transpositions(m1, m2)) / 2
	l1 := float64(len(s1))
	l2 := float64(len(s2))

	return (1.0 / 3) * (lm1/l1 + lm2/l2 + (lm1-t)/lm1)
}

func commonCharacters(s1, s2 string) []byte {
	// output slice
	common := make([]byte, 0)
	// calculate the maximum distance within which a character is common
	maxDistance := int(math.Floor(math.Min(float64(len(s1)), float64(len(s2))) / 2))

	// loop through all characters of s1
	for i := 0; i < len(s1); i++ {
		char := s1[i]
		// check if the current character is common in s2
		charCommon := charCommon(char, i, s2, maxDistance)
		// if it is common append it to the output
		if charCommon {
			common = append(common, char)
		}
	}
	return common
}

func charCommon(char byte, pos int, otherString string, maxDistance int) bool {
	for i := 0; i < len(otherString); i++ {
		c := otherString[i]
		// not the same character
		if char != c {
			continue
		}
		// calculate the distance between the char and the current character c
		if dist := int(math.Abs(float64(pos - i))); dist <= maxDistance {
			return true
		}
	}
	return false
}

func transpositions(a, b []byte) (t int) {
	la := len(a)
	lb := len(b)

	// internalTranspositions counts the number of transpositions
	// it requires that a is smaller than b
	internalTranspositions := func(a, b []byte) (t int) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				t++
			}
		}
		return
	}

	// check for the smaller slice of bytes and call internalTranspositions accordingly
	if la <= lb {
		t = internalTranspositions(a, b)
	} else {
		t = internalTranspositions(b, a)
	}
	// add all the transpositions where a and b are not of equal length
	t += int(math.Abs(float64(la) - float64(lb)))
	return
}

//jaroWinkler////////////////////////////////////////////////////////////////
func JaroWinkler(s1, s2 string) (distance float64) {
	// set distance to JaroScore
	distance = Jaro(s1, s2)
	// check if the distance is above the boost threshold and therefor the boost is applied
	if distance >= boostThreshold {
		// add to the distance a factor which includes the prefix length a scale for this length and a factor that depends on the jaro distance
		distance = distance + (float64(prefixLength(s1, s2)) * prefixScale * (1 - distance))
	}
	return
}

func SetBoostThreshold(bt float64) error {
	if bt < 0 || bt > 1 {
		return errors.New("Bt not in correct range [0.0,1.0]")
	}
	boostThreshold = bt
	return nil
}

func SetPrefixScale(p float64) error {
	if p <= 0 || p >= 0.25 {
		return errors.New("p not in correct range ]0.0,1.0]")
	}
	prefixScale = p
	return nil
}

func prefixLength(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	var size int

	// calculate size for loop
	// take the length of the smaller string as the prefix can have a maximum of the length of the smaller string
	if l1 >= l2 {
		size = l2
	} else {
		size = l1
	}

	l := 0
	for i := 0; i < size; i++ {
		e1 := s1[i]
		e2 := s2[i]
		// break condition when either the maximum prefix length is reached or the two characters do not match
		if l == MAXPREFIXLENGTH || e1 != e2 {
			break
		}
		l++
	}
	return l
}

//Interal////////////////////////////////////////////////////////
func snaker(s string, wordSeparator rune, firstRune func(rune) rune, firstRuneOfWord func(rune) rune, otherRunes func(rune) rune) string {
	useWordSeperator := wordSeparator != rune(0)
	newS := []rune{}

	// pops a rune off newS
	lastRuneIsWordSeparator := func() bool {
		if len(newS) > 0 {
			return newS[len(newS)-1] == wordSeparator
		}
		return true
	}

	prev := wordSeparator
	for _, cur := range s {
		isWordBoundary := (unicode.IsLower(prev) && unicode.IsUpper(cur)) || unicode.IsSpace(prev)

		if !unicode.IsLetter(cur) {
			// ignore
		} else if isWordBoundary {
			if useWordSeperator && !lastRuneIsWordSeparator() {
				newS = append(newS, wordSeparator)
			}
			newS = append(newS, firstRuneOfWord(cur))
		} else {
			newS = append(newS, otherRunes(cur))
		}

		prev = cur
	}

	if len(newS) > 0 {
		newS[0] = firstRune(newS[0])
	}

	return string(newS)
}

//String Format////////////////////////////////////////////////
func Format(formatStr string, a ...interface{}) string {
	split := strings.SplitAfter(formatStr, "{}")
	for i, arg := range a {
		split[i] = formatPiece(split[i], arg)
	}
	return strings.Join(split, "")
}

func formatPiece(piece string, arg interface{}) string {
	return strings.Replace(piece, "{}", as.ToString(arg), 1)
}

//String Slice////////////////////////////////////////////////
func Add(data []string, otherList []string) []string {
	n := len(data) + len(otherList)
	newList := make([]string, n, n)
	for _, v := range data {
		if !(Includes(data, v)) {
			newList = append(newList, v)
		}
	}
	for _, v := range otherList {
		if !(Includes(data, v)) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Subtract(data []string, excludeList []string) []string {
	newList := make([]string, 0, len(data))
	for _, v := range data {
		if !(Includes(excludeList, v)) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Includes(data []string, arg string) bool {
	result := false
	for _, v := range data {
		if v == arg {
			result = true
			break
		}
	}
	return result
}

func Equals(data []string, otherList []string) bool {
	l1 := Subtract(data, otherList)
	l2 := Subtract(otherList, data)
	l3 := Add(l1, l2)
	count := 0
	for _, v := range l3 {
		if v != "" {
			count += 1
		}
	}
	return (count == 0)
}

//Ordinise///////////////////////////////////////////////////////
func Ordinise(number int) string {
	switch int(math.Abs(float64(number))) % 100 {
	case 11, 12, 13:
		return "th"
	default:
		switch int(math.Abs(float64(number))) % 10 {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		}
	}
	return "th"
}

func ToOrdinise(number int) string {

	return as.ToString(number) + Ordinise(number)
}
