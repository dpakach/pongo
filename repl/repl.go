package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/dpakach/pongo/lexer"
	"github.com/dpakach/pongo/parser"
	"github.com/dpakach/pongo/evaluator"
	"github.com/dpakach/pongo/object"
)

const PROMPT = ":-/ >> "

var stack []string

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
		stack = append(stack, line)
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
