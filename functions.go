package functions

import (
	"html/template"
	"os"
	"strings"
)

var (
	Bold = func(content string) (template.HTML, error) {
		return template.HTML("<b>" + content + "</b>"), nil
	}
	Italic = func(content string) (template.HTML, error) {
		return template.HTML("<i>" + content + "</i>"), nil
	}
	Underline = func(content string) (template.HTML, error) {
		return template.HTML("<u>" + content + "</u>"), nil
	}
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetFilenameNoExtension(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n >= 0 {
		return s[:n]
	}
	return s
}

func RemoveDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
