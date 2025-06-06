package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		io.WriteString(out, "Parser:")
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		output := evaluator.Eval(program)

		io.WriteString(out, "Eval:")
		io.WriteString(out, output.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

const MONKEY_FACE = `
	.--. .-"	"-. .--.
   / .. \/ .-. .-. \/ .. \
     | | '| / Y \ |' ||
    | \ \ \ 0 | 0 / / /|
  \ '- ,\.-"""""""-./,-' /
     ''-' /_ ^ ^ _\ '-''
         | \._ _./ |
         \ \ '~' / /
        '._ '-=-' _.'
           '-----'`
