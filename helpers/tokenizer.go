package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

type Token struct {
	Type  string
	Value string
}

func Tokenize(rule string) []Token {
	re := regexp.MustCompile(`\s*(\(|\)|AND|OR|>|<|=|'[^']*'|\w+)\s*`)
	matches := re.FindAllStringSubmatch(rule, -1)
	var tokens []Token

	for _, match := range matches {
		tokenValue := match[1]
		tokenType := ""
		switch tokenValue {
		case "AND", "OR", ">", "<", "=":
			tokenType = "Operator"
		case "(", ")":
			tokenType = "Parenthesis"
		default:
			if strings.HasPrefix(tokenValue, "'") && strings.HasSuffix(tokenValue, "'") {
				tokenType = "Literal"
			} else if _, err := strconv.Atoi(tokenValue); err == nil {
				tokenType = "Literal"
			} else {
				tokenType = "Identifier"
			}
		}
		tokens = append(tokens, Token{Type: tokenType, Value: tokenValue})
	}
	return tokens

}
