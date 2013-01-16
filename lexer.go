package goawesome

import (
	"regexp"
)

type token map[string]interface{}

func tokenizer(code []byte) []token {
	var tokens []token
	numberRegex, _ := regexp.Compile("([0-9]+)")
	if number := numberRegex.Find(code); len(number) > 0 {
		tokens = append(tokens, token{"type": "NUMBER", "value": string(number)})
	}
	return tokens
}
