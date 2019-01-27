package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/anraku/monkey/evaluator"
	"github.com/anraku/monkey/lexer"
	"github.com/anraku/monkey/object"
	"github.com/anraku/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator := evaluator.Eval(program, env)

		if evaluator != nil {
			io.WriteString(out, evaluator.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
