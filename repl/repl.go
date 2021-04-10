package repl

import (
	"bufio"
	"fmt"
	"github.com/dpakach/pongo/compiler"
	"github.com/dpakach/pongo/lexer"
	"github.com/dpakach/pongo/parser"
	"github.com/dpakach/pongo/vm"
	"io"
)

const PROMPT = ":-/ >> "

var stack []string

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "woops! compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! executing bytecode failed:\n %s\n", err)
		}

		lastPopped := machine.LastPoppedElem()

		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
