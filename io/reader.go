package io

import (
	"os"
	"thl/colors"
	"thl/functions"
	"thl/types"

	"fmt"
	"strings"

	. "thl/constants"

	"github.com/mattn/go-tty"
	"golang.org/x/term"
)

// GetBool prompts the user with a Yes/No question and returns true for "Y" and false for "N".
func GetBool(message string, color types.IColor) bool {
	// Open TTY
	tty, err := tty.Open()
	if err != nil {
		colors.Red.Println(err)
	}
	defer tty.Close()

	// Switch to raw mode
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)

	if err != nil {
		colors.Red.Println(err)
	}

	defer term.Restore(fd, oldState)

	// Print message with color
	color.Print(message, " (Y/n) ")

	// Read a single key press
	char, err := tty.ReadRune()
	if err != nil {
		colors.Red.Println(err)
	}

	// Print a newline for cleaner output
	fmt.Print("\n")
	fmt.Print("\n")
	functions.ClearCurrentLine()

	// Normalize input
	return char == '\r' || char == '\n' || char == 'Y' || char == 'y'
}

func GetRules(terminals, nonTerminals []string, axiom string) []types.Rule {
	var rules []types.Rule

	for {
		rule := ReadRule(nonTerminals, terminals, rules, colors.Magenta)
		if rule != nil {
			rules = append(rules, *rule)
		} else {
			if ValidateRules(rules, axiom, terminals) {
				// if user entered valid rules, we
				// can break the loop and return the rules!
				break
			}
		}
	}

	return rules
}

func ValidateRules(rules []types.Rule, axiom string, terminals []string) bool {
	var atLeastOneIn bool = false
	var allRightTermOut bool = false

	for i := 0; i < len(rules); i++ {
		r := rules[i]
		if len(r.Left) == 2 && r.Left[0] == axiom && r.Left[1] == REAL_NT_SUFFIX {
			atLeastOneIn = true
		}

		every := true
		for j := 0; j < len(r.Right); j++ {
			if r.Right[j] == REAL_NT_SUFFIX {
				every = false
			}
		}

		if every {
			allRightTermOut = true
		}
	}
	if len(rules) == 0 {
		colors.Red.Println("✗ At least one rule required to continue...")
		return false
	} else if !atLeastOneIn {
		colors.Red.Println("✗ Missing one rule with axiom in left side")
		return false
	} else if !allRightTermOut {
		colors.Red.Println("✗ Missing one rule with all terminals in right side")
		return false
	}
	return true
}

func GetInput(message string, symbol string, color types.IColor) []string {
	var items []string
	var count int = 1

	color.Print(message, "\n")
	for {
		color.Print(STR_WALL)

		var input string
		colors.Cyan.Printf("%s[%d]: ", symbol, count)
		fmt.Scanf("%s", &input)

		if len(input) > 0 {
			input = string(strings.TrimSpace(input)[0])
		}

		color.Print(STR_WALL)
		colors.Cyan.Printf("%s[%d]: ", symbol, count)
		fmt.Printf("%s", input)

		count++

		if input == "" && len(items) == 0 {
			functions.ClearCurrentLine()
			color.Print(STR_WALL)
			colors.Red.Println("✗ Please enter at least one item!")
			count--
			continue
		} else if input == "" {
			functions.ClearCurrentLine()
			count--
			break
		}

		if functions.InArray(items, input) {
			functions.ClearCurrentLine()
			color.Print(STR_WALL)
			colors.Red.Println("✗ Duplicate entry!")
			count--
			continue
		}

		functions.ClearCurrentLine()
		color.Print(STR_WALL)
		colors.Cyan.Printf("%s[%d]: ", symbol, count-1)
		fmt.Println(input)

		items = append(items, input)
	}

	color.Println(strings.Repeat("-", len(message)-2), "\n")
	return items
}

func GetAxiom(nonTerminals []string, color types.IColor) string {
	const message string = "Axiom S (↩ finish)"
	var axiom string

	color.Print(message, "\n")

	for {
		color.Print(STR_WALL)
		color.Print("Axiom (ex. S): ")

		fmt.Scanf("%s", &axiom)

		if functions.InArray(nonTerminals, axiom) {
			break
		}

		functions.ClearLine(1)

		color.Print(STR_WALL)
		colors.Red.Print("✗ Axiom must be in Non-terminals!\n")
	}

	color.Println(strings.Repeat("-", len(message)), "\n")

	return axiom
}

func GetString(message string, color types.IColor) string {
	var input string

	color.Print(message)
	fmt.Scanf("%s", &input)

	return input
}

func GetInt(message string, color types.IColor) int {
	var input int

	color.Print(message)
	fmt.Scanf("%d", &input)

	return input
}
