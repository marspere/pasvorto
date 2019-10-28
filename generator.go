package pasvorto

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Password struct {
	Value  string
	Length int
	// Regular : indicate password rules, allowing the
	// flowing four values([0, 1, 2, 3]).
	Regular []int
	// Score : [0,1,2,3,4] if crack time is less than
	// [10^2, 10^4, 10^6, 10^8, Infinity].
	// (useful for implementing a strength bar.)
	Score int
}

const (
	Letter = iota
	Digital
	Underline
	Characters
)

const (
	letter     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	digital    = "0123456789"
	underline  = `_`
	characters = `*&$%#@`
)

// Generate used to generate a random password based on the parameters,
// the default length of password is 8, the default password rule
// is composed of numbers and letters.
func Generate(length int, regular ...int) *Password {
	pwd := newPwd(length, regular)
	pwd.generate()
	pwd.scoring()
	return pwd
}

func newPwd(length int, regular []int) *Password {
	return &Password{
		Length:  length,
		Regular: regular,
	}
}

func (pwd *Password) generate() {
	tempPwd := make([]byte, 0)
	tempPwd = append(tempPwd, letter[rand.Intn(len(letter))])
	flag := contains(pwd.Regular)
	if pwd.Length < 8 {
		pwd.Length = 8
	}
	if flag == 6 {
		tempPwd = append(tempPwd, []byte(underline)...)
		tempPwd = append(tempPwd, characters[rand.Intn(len(characters))])
		for i := 0; i < pwd.Length-3; i++ {
			tempPwd = append(tempPwd, (letter + digital)[rand.Intn(len(letter + digital))])
		}
	} else if flag == 3 {
		tempPwd = append(tempPwd, []byte(underline)...)
		for i := 0; i < pwd.Length-2; i++ {
			tempPwd = append(tempPwd, (letter + digital)[rand.Intn(len(letter + digital))])
		}
	} else if flag == 4 {
		tempPwd = append(tempPwd, characters[rand.Intn(len(characters))])
		for i := 0; i < pwd.Length-2; i++ {
			tempPwd = append(tempPwd, (letter + digital)[rand.Intn(len(letter + digital))])
		}
	} else {
		for i := 0; i < pwd.Length-1; i++ {
			tempPwd = append(tempPwd, (letter + digital)[rand.Intn(len(letter+digital))])
		}
	}
	endPwd := shuffle(tempPwd[1:])
	tempPwd = append(tempPwd, tempPwd[0])
	tempPwd = append(tempPwd, endPwd...)
	pwd.Value = string(tempPwd[pwd.Length:])
	return
}

func shuffle(regular []byte) []byte {
	rand.Shuffle(len(regular), func(i, j int) {
		regular[i], regular[j] = regular[j], regular[i]
	})
	return regular
}

func contains(regFlag []int) (flag int) {
	for _, value := range regFlag {
		flag += value
	}
	return
}
