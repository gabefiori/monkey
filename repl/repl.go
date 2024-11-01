package repl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	reader := bytes.NewReader(scanner.Bytes())
	l := lexer.New(reader)

	for {
		fmt.Fprint(w, PROMPT)

		if !scanner.Scan() {
			return
		}

		reader.Reset(scanner.Bytes())
		l.Prepare()

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tok)
		}
	}
}
