package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

type TokenType string

const (
	Operator    TokenType = "operator"
	Operand     TokenType = "operand"
	Parenthesis TokenType = "parenthesis"
)

type Token struct {
	Type  TokenType
	Value string
}

func TokenizeRule(rule string) []Token {
	tokens := []Token{}
	var currentVal strings.Builder

	operatorRegex := regexp.MustCompile(`<=|>=|<|>|=`)

	i := 0
	for i < len(rule) {
		ch := rule[i]

		if unicode.IsSpace(rune(ch)) {
			i += 1
			continue
		}

		if ch == '(' || ch == ')' {
			if currentVal.Len() > 0 {
				tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
				currentVal.Reset()
			}
			tokens = append(tokens, Token{Type: Parenthesis, Value: string(ch)})
			i += 1
		} else if strings.HasPrefix(rule[i:], "AND") || strings.HasPrefix(rule[i:], "OR") {
			if currentVal.Len() > 0 {
				tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
				currentVal.Reset()
			}
			if strings.HasPrefix(rule[i:], "AND") {
				tokens = append(tokens, Token{Type: Operator, Value: "AND"})
				i += 3
			}
			if strings.HasPrefix(rule[i:], "OR") {
				tokens = append(tokens, Token{Type: Operator, Value: "OR"})
				i += 2
			}
		} else if i+1 < len(rule) && operatorRegex.MatchString(rule[i:i+2]) {
			if currentVal.Len() > 0 {
				tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
				currentVal.Reset()
			}
			tokens = append(tokens, Token{Type: Operator, Value: rule[i : i+2]})
			i += 2
		} else if operatorRegex.MatchString(string(ch)) {
			if currentVal.Len() > 0 {
				tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
				currentVal.Reset()
			}
			tokens = append(tokens, Token{Type: Operator, Value: string(ch)})
			i += 1
		} else if ch == '\'' {
			if currentVal.Len() > 0 {
				tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
				currentVal.Reset()
			}
			i += 1
			for i < len(rule) && rule[i] != '\'' {
				currentVal.WriteRune(rune(rule[i]))
				i += 1
			}
			tokens = append(tokens, Token{Type: Operand, Value: currentVal.String()})
			currentVal.Reset()
			i += 1
		} else {
			currentVal.WriteRune(rune(rule[i]))
			i += 1
		}
	}

	return tokens
}
