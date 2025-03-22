package io

import (
	"fmt"
	"strings"
	"thl/colors"
	"thl/functions"
	"thl/types"
)

// ReadRule reads a grammar rule from the user input, validates it, and returns it.
func ReadRule(nonterm []string, term []string, rules []types.Rule, c types.IColor) *types.Rule {
	var num int = len(rules)
	var left string = ""
	var right string = ""
	var isLValid bool = false
	var isRValid bool = false
	var exists bool = false

	const leftTitle string = "Rule's Left (↩ continue/cancel)\n| "

	for !isLValid {
		left = ""

		if num == 0 {
			c.Print(leftFirstTitle)
		} else {
			c.Print(leftTitle)
		}

		// gap between terminal and last line
		functions.Gap()
		c.Print(wall)

		fmt.Scanf("%s", &left)
		left = functions.Trim(left)

		// if user pressed enter, we can cancel the rule!
		// return nil to cancel the rule.
		if len(left) == 0 {
			functions.ClearLine(2)
			return nil
		}

		// continue and validate the input entered...
		// if valid, the loop stops and we can continue
		// to the right side of the rule.
		isLValid = ValidateInput(left, term, nonterm)

		if !isLValid {
			continue
		}

	}

	for !isRValid {
		right = ""

		// clear the left side of the rule
		// and display the prompt again...

		functions.ClearLine(2)
		c.Print(rightTitle)
		fmt.Printf("%s %s ", left, separator)

		// scan the right side of the rule and validate it.
		fmt.Scanf("%s", &right)
		right = functions.Trim(right)

		isRValid = ValidateInput(right, term, nonterm)

		// if the right side of the rule is invalid,
		// display the invalid rule and continue the loop.
		if !isRValid {
			continue
		}
	}

	// check if the rule already exists in the rules list.
	// if it does, display a message and continue the loop.

	exists = false
	for i := 0; i < len(rules); i++ {
		r := rules[i]
		if functions.Join(r.Left) == left && functions.Join(r.Right) == right {
			exists = true
			break
		}
	}
	if exists {
		functions.ClearLine(2)
		colors.Red.Print("Rule already exists.\n")
		return ReadRule(nonterm, term, rules, c)
	}

	// clear the prompt message and display the rule
	// as a result of the user's input _____________

	functions.Clear()

	c.Printf("Rule (Number %d)\n| ", num+1)
	fmt.Printf("%s %s %s\n", left, separator, right)
	c.Print("-------------------\n\n")
	// ____________________________________________

	// return the rule as a pointer to the rule.

	return &types.Rule{
		Left:  functions.Split(left),
		Right: functions.Split(right),
	}
}

// ValidateInput checks if each character in the input string is present in either
// the term or nonterm slices. If any character is not found in either slice, it
// clears the Rule prompt and returns false, indicating invalid input. Otherwise,
// it returns true.
func ValidateInput(input string, term []string, nonterm []string) bool {
	var isNonTerm bool = false

	for i := 0; i < len(input); i++ {
		char := string(input[i])

		if char == NT_SUFFIX {
			// error, unexpected symbol
			functions.ClearLine(2)
			ErrorPos(input, i, "Unexpected '"+NT_SUFFIX+"'")
			return false
		}

		if i < len(input)-1 && string(input[i+1]) == NT_SUFFIX {
			isNonTerm = true
			i++
		} else {
			isNonTerm = false
		}

		if isNonTerm && !functions.InArray(nonterm, char) {

			functions.ClearLine(2)
			ErrorPos(input, i-1, "Symbol not found in non-terminals")
			return false

		} else if !isNonTerm && !functions.InArray(term, char) {

			functions.ClearLine(2)
			ErrorPos(input, i, "Symbol not found in terminals")
			return false
		}

	}

	return true
}

func ErrorPos(str string, pos int, message string) {

	position := pos - 1

	if pos == 0 {
		position = 0
	}

	colors.Red.Printf("%s\n", str)
	colors.Red.Printf("%s^\n", strings.Repeat(" ", pos))
	colors.Gray.Printf("%s%s\n", strings.Repeat(" ", position), message)

	fmt.Println()
	fmt.Println()
}
