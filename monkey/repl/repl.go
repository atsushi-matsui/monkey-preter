package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/atsushi-matsui/monkey-preter/monkey/lexer"
	"github.com/atsushi-matsui/monkey-preter/monkey/token"
)

const PROMPT  = ">> "


func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		fmt.Printf(PROMPT)

		l := lexer.New(scanner.Text())
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
