package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

const MAX_ASCII = 127

var sensitiveWords = []string{"password", "api_key", "secret", "token"}

func check(pass *analysis.Pass, call *ast.CallExpr) {
	arg := call.Args[0]
	rawMsg, ok := arg.(*ast.BasicLit)
	if !ok || rawMsg.Kind != token.STRING {
		return
	}

	// AST не может распарсить строку не правильно
	// поэтому можно не отлавливать ошибку от unquote
	msg, _ := strconv.Unquote(rawMsg.Value)

	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsUpper(r) {
		pass.Reportf(arg.Pos(), "log should be in lowercase")
	}

	var (
		foundNonEnglish bool
		foundSpecial    bool
	)
	for _, r := range msg {
		if !foundNonEnglish && r > MAX_ASCII {
			foundNonEnglish = true
		}
		if !foundSpecial && !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			foundSpecial = true
		}
		if foundNonEnglish && foundSpecial {
			break
		}
	}

	if foundNonEnglish {
		pass.Reportf(arg.Pos(), "log message should be in english")
	}
	if foundSpecial {
		pass.Reportf(arg.Pos(), "log message should not contain emojis or special symbols")
	}

	msgLower := strings.ToLower(msg)
	for _, word := range sensitiveWords {
		if strings.Contains(msgLower, word) {
			pass.Reportf(arg.Pos(), "log contains sensitive data: %s", word)
			return
		}
	}
}
