package terminal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	YesDefintions     = []string{"y", "Y", "yes", "YES", "Yes"}
	NoDefinitions     = []string{"n", "N", "no", "NO", "No"}
	YesOrNoSuffix     = "? [Y/N]: "
	ErrProvideYesOrNo = errors.New("please answer with 'yes' or 'no'")
	ErrInvalidChoice  = errors.New("Invalid choice please type a number in range")
)

var Writer io.Writer = os.Stdout

type Dialog struct {
}

//YesOrNo ask a question that can be answered by yes or no
func YesOrNo(question string, def bool) bool {
	//ask
	fmt.Fprintf(Writer, question+YesOrNoSuffix)

	//answer
	s, _ := readline()
	if hasString(YesDefintions, s) {
		return true
	} else if hasString(NoDefinitions, s) {
		return false
	} else if s == "" {
		return def
	} else {
		fmt.Fprintln(Writer, ErrProvideYesOrNo.Error())
		return YesOrNo(question, def)
	}
}

//AskString asks for a string
func AskString(question string) string {
	fmt.Fprintf(Writer, "%s:\n\t", question)
	s, _ := readline()
	return s
}

func Choice(question string, options []string) int {
	fmt.Fprintf(Writer, "%s:\n", question)
	for i, q := range options {
		fmt.Fprintf(Writer, "%d) %s\n", (i + 1), q)
	}
	s, _ := readline()
	answer, _ := strconv.Atoi(s)

	if answer == 0 || answer > len(options) {
		fmt.Fprintln(Writer, ErrInvalidChoice.Error())
		return Choice(question, options)
	}

	//in programming slices start at 0 not one
	return answer - 1
}

//readline reads input from the user byte by byte
func readline() (value string, err error) {
	var valb []byte
	var n int
	b := make([]byte, 1)
	for {
		// read one byte at a time so we don't accidentally read extra bytes
		n, err = os.Stdin.Read(b)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 || b[0] == '\n' {
			break
		}
		valb = append(valb, b[0])
	}

	return strings.TrimSuffix(string(valb), "\r"), nil
}

//indexOfString returns the index a string in a string slice, returns -1
//if the given string is not found
func indexOfString(h []string, n string) int {
	for i, v := range h {
		if v == n {
			return i
		}
	}

	return -1
}

//hasString is a wrapper around indexOfString
func hasString(h []string, n string) bool {
	if indexOfString(h, n) != -1 {
		return true
	}
	return false
}
