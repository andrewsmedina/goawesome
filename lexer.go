package goawesome

import (
	"regexp"
)

type token map[string]interface{}

func tokenizer(code []byte) []token {
	var tokens []token
	numberRegex, _ := regexp.Compile("([0-9]+)")
	operatorRegex, _ := regexp.Compile(`(\|\||&&|==|!=|<=|>=)`)
	if number := numberRegex.Find(code); len(number) > 0 {
		tokens = append(tokens, token{"type": "NUMBER", "value": string(number)})
	} else if operator := string(operatorRegex.Find(code)); len(operator) > 0 {
		tokens = append(tokens, token{"type": operator, "value": operator})
	} else {
		t := string(code)
		tokens = append(tokens, token{"type": t, "value": t})
	}
	return tokens
}
