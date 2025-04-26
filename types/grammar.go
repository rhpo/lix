package types

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	. "thl/constants"
	. "thl/functions"
)

type GrammarFile struct {
	T []string `json:"T"`
	N []string `json:"N"`
	S string   `json:"S"`
	P []Rule   `json:"P"`
}

type Grammar struct {
	T []string
	N []string
	S string
	P []Rule
}

func CreateGrammar(T []string, N []string, S string, P []Rule) Grammar {
	return Grammar{
		T: T,
		N: N,
		S: S,
		P: P,
	}
}

// The Chomsky Hierarchy classifies formal grammars into four types based on their production rules and computational power:
func (g *Grammar) Type() string {

	if g.isType3() {
		return "Type 3: Regular Grammar"
	} else if g.isType2() {
		return "Type 2: Context-Free Grammar"
	} else if g.isType1() {
		return "Type 1: Context-Sensitive Grammar"
	} else if g.isType0() {
		return "Type 0: Unrestricted Grammar"
	}

	return "Unknown Type"
}

func (g *Grammar) isType3() bool {

	for i := 0; i < len(g.P); i++ {
		rule := g.P[i]
		if len(rule.Left) != 2 || rule.Left[1] != REAL_NT_SUFFIX || !InArray(g.N, rule.Left[0]) || !IsRuleOnlyTerminal(rule, g.T) && !IsRuleRightRegular(rule, g.N, g.T) && !IsRuleLeftRegular(rule, g.N, g.T) {
			return false
		}
	}

	return true
}

func (g *Grammar) isType2() bool {

	for i := 0; i < len(g.P); i++ {
		R := g.P[i]
		if len(R.Left) != 2 || R.Left[1] != REAL_NT_SUFFIX || !InArray(g.N, R.Left[0]) {
			return false
		}
		if len(R.Right) == 0 {
			return false
		}
	}

	return true
}

func (g *Grammar) isType1() bool {

	for i := 0; i < len(g.P); i++ {
		R := g.P[i]

		if !IsRuleTerminalNonterminal(R, g.N, g.T) || !LeftSideGamma(R, g.N, g.T) {
			return false
		}

	}

	return true
}

func (g *Grammar) isType0() bool {
	return true
}

func (g *Grammar) SaveTo(name string) {
	// check for the existence of a local folder called "grammars"
	dir := "grammars"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("Error creating direct")
			return
		}
	}

	// create a GrammarJson object
	grammarJson := GrammarFile{
		T: g.T,
		N: g.N,
		S: g.S,
		P: g.P,
	}

	// convert the Grammar object to a GrammarJson object
	data, err := json.MarshalIndent(grammarJson, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling grammar to J")
		return
	}

	// save the GrammarJson object to a file
	filePath := filepath.Join(dir, name+".gm")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating f")
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to f")
		return
	}

	fmt.Println("Grammar saved to file:", filePath)
}

func (g *Grammar) OpenFrom(name string) error {
	// check for the existence of a local folder called "grammars"

	dir := "grammars"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory")
		}
	}

	// open the file
	filePath := filepath.Join(dir, name+".gm")
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file")
	}
	defer file.Close()

	// decode the file
	decoder := json.NewDecoder(file)
	grammarJson := GrammarFile{}
	err = decoder.Decode(&grammarJson)
	if err != nil {
		return fmt.Errorf("error decoding file")
	}

	// convert the GrammarJson object to a Grammar object
	g.T = grammarJson.T
	g.N = grammarJson.N
	g.S = grammarJson.S
	g.P = grammarJson.P

	return nil
}
