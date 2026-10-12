package main

import (
	"os"
	"strconv"
	"strings"
)

func capitalize(s string) string {
	if s == "" {
		return ""
	}

	f := strings.ToUpper(s[:1])
	l := strings.ToLower(s[1:])
	s = f + l

	return s
}

func applyLastN(words []string, n int, command string) []string {
	start := len(words) - n

	if n > len(words) {
		start = 0
	}

	for i := start; i < len(words); i++ {
		if command == "up" {
			words[i] = strings.ToUpper(words[i])
		} else if command == "low" {
			words[i] = strings.ToLower(words[i])
		} else if command == "cap" {
			words[i] = capitalize(words[i])
		}
	}

	return words
}

func executeCommands(words []string) []string {
	var res []string

	for i := 0; i < len(words); {
		if words[i] == "(up)" {
			if len(res) > 0 {
				res[len(res)-1] = strings.ToUpper(res[len(res)-1])
			}
			i++
			continue
		} else if words[i] == "(low)" {
			if len(res) > 0 {
				res[len(res)-1] = strings.ToLower(res[len(res)-1])
			}
			i++
			continue
		} else if words[i] == "(cap)" {
			if len(res) > 0 {
				res[len(res)-1] = capitalize(res[len(res)-1])
			}
			i++
			continue
		}

		if words[i] == "(hex)" {
			if len(res) > 0 {
				val, err := strconv.ParseInt(res[len(res)-1], 16, 64)
				if err == nil {
					res[len(res)-1] = strconv.FormatInt(val, 10)
				}
			}
			i++
			continue
		} else if words[i] == "(bin)" {
			if len(res) > 0 {
				val, err := strconv.ParseInt(res[len(res)-1], 2, 64)
				if err == nil {
					res[len(res)-1] = strconv.FormatInt(val, 10)
				}
			}
			i++
			continue
		}

		if words[i] == "(up," || words[i] == "(low," || words[i] == "(cap," {
			if i+1 < len(words) {
				txtNum := strings.TrimSuffix(words[i+1], ")")
				n, err := strconv.Atoi(txtNum)

				if err == nil {
					if words[i] == "(up," {
						applyLastN(res, n, "up")
					} else if words[i] == "(low," {
						applyLastN(res, n, "low")
					} else if words[i] == "(cap," {
						applyLastN(res, n, "cap")
					}

					i += 2
					continue
				}
			}
		}

		res = append(res, words[i])
		i++
	}

	return res
}

func isPunc(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
		return true
	}

	return false
}

func separatePunctuation(words []string) []string {
	var res []string

	for _, word := range words {
		var current string
		var punc string

		for _, r := range word {
			if isPunc(r) {
				if current != "" {
					res = append(res, current)
					current = ""
				}

				punc += string(r)
			} else {
				if punc != "" {
					res = append(res, punc)
					punc = ""
				}

				current += string(r)
			}
		}

		if current != "" {
			res = append(res, current)
		}

		if punc != "" {
			res = append(res, punc)
		}
	}

	return res
}

func isPuncToken(token string) bool {
	if token == "" {
		return false
	}

	for _, r := range token {
		if !isPunc(r) && r != '\'' {
			return false
		}
	}

	return true
}

func rebuildText(words []string) string {
	var res string

	for _, word := range words {
		if isPuncToken(word) {
			res += word
		} else {
			if res == "" {
				res += word
			} else {
				res += " " + word
			}
		}
	}

	return res
}

func quote(words []string) []string {
	var res []string
	inQuote := false
	openQuote := false

	for _, word := range words {
		if word == "'" {
			if !inQuote {
				inQuote = true
				openQuote = true
			} else {
				if len(res) > 0 {
					res[len(res)-1] += "'"
				}

				inQuote = false
				openQuote = false
			}

			continue
		}

		if openQuote {
			word = "'" + word
			openQuote = false
		}

		res = append(res, word)
	}

	return res
}

func fixAn(words []string) []string {
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			next := words[i+1]

			if next != "" {
				first := strings.ToLower(next[:1])

				if first == "a" ||
					first == "e" ||
					first == "i" ||
					first == "o" ||
					first == "u" ||
					first == "h" {

					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}

	return words
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}

	text := string(data)
	words := strings.Fields(text)

	words = executeCommands(words)
	words = separatePunctuation(words)
	words = fixAn(words)
	words = quote(words)

	finalText := rebuildText(words)

	err = os.WriteFile(outputFile, []byte(finalText), 0644)
	if err != nil {
		return
	}
}
