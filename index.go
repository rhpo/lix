package main

import (
	"thl/colors"
	"thl/eval"
	"thl/functions"
	"thl/types"

	"thl/io"
)

var (
	TerminalColor    *colors.ColorType = colors.Green
	NonTerminalColor *colors.ColorType = colors.Blue
	AxiomColor       *colors.ColorType = colors.Yellow
	RuleColor        *colors.ColorType = colors.Magenta
)

var G types.Grammar

func ReadGrammar() {
	G.T = io.GetInput("Terminals T (↩ finish)", "T", TerminalColor)
	G.N = io.GetInput("Non-terminals N (↩ finish)", "N", NonTerminalColor)
	G.S = io.GetAxiom(G.N, AxiomColor)

	functions.Clear()
	G.P = io.GetRules(G.T, G.N, G.S)

	ShowGrammar()
}

func ShowGrammar() {

	colors.Gray.Println(G.Type())

	io.PrintTable(G.T, "\nTerminals", TerminalColor)
	io.PrintTable(G.N, "Non-terminals", NonTerminalColor, G.S)
	io.PrintTable([]string{G.S}, "Axiom: ", AxiomColor)

	if len(G.P) > 0 {
		// passs G.N to highlight non-terminals in the rules
		io.PrintRules(G.P, G.N, RuleColor)
	}

	println()
}

func app(open bool) {
	functions.Help()

	// whether we should open grammar from a file, or create a new one
	if open {
		name := io.GetString("Name? ", colors.Blue)
		error := G.OpenFrom(name)

		if error != nil {
			colors.Red.Println(error)
			return
		}

		ShowGrammar()
	} else {

		ReadGrammar()
		functions.Clear()
		ShowGrammar()

		wlen := io.GetInt("Words maximum? ", colors.Yellow)
		if wlen < 1 {
			colors.Red.Println("✗ Words count must be greater than 0!")
			return
		}

		var words []string = eval.Evaluate(G, 0, wlen)

		if len(words) == 0 {
			colors.Red.Println("✗ No words generated!")
			return
		}

		io.PrintTable(words, "\nGenerated Words (w)", colors.Cyan)

		if io.GetBool("Save grammar to file?", colors.Yellow) {
			name := io.GetString("Name? ", colors.Blue)
			G.SaveTo(name)
		}
	}
}

func main() {
	functions.Help()

	for {
		open := io.GetBool("Open grammar from file?", colors.Yellow)

		app(open)
	}
}
