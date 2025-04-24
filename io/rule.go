package io

import (
	"fmt"
	"strings"
	"thl/colors"
	"thl/functions"
	"thl/types"

	. "thl/constants"
)

// ReadRule reads a grammar rule from the user input, validates it, and returns it.
func ReadRule(nonterm []string, term []string, rules []types.Rule, c types.IColor) *types.Rule {
	var num int = len(rules)
	var left string = ""
	var right string = ""
	var isLValid bool = false
	var isRValid bool = false
	var exists bool = false

	const STR_TITLE_LEFT string = "Rule's Left (↩ continue/cancel)\n| "

	for !isLValid {
		left = ""

		if num == 0 {
			c.Print(STR_TITLE_LEFT_FIRST)
		} else {
			c.Print(STR_TITLE_LEFT)
		}

		// gap between terminal and last line
		functions.Gap()
		c.Print(STR_WALL)

		fmt.Scanf("%s", &left)
		left = functions.Trim(left)

		// if user pressed enter, we can cancel the rule!
		// return nil to cancel the rule.
		if len(left) == 0 {
			functions.ClearLine(2)
			return nil
		}

		if reformat(left) != nil {
			left = *reformat(left)
		} else {
			functions.ClearLine(2)
			ErrorPos(left, len(left), "Expected character after "+ESCAPE_CHARACTER)
			continue
		}

		// we now have the input

		// continue and validate the input entered...
		// if valid, the loop stops and we can continue
		// to the right side of the rule.
		isLValid = ValidateInput(&left, term, nonterm)

		if !isLValid {
			continue
		}

	}

	for !isRValid {
		right = ""

		// clear the left side of the rule
		// and display the prompt again...

		functions.ClearLine(2)
		c.Print(STR_TITLE_RIGHT)
		fmt.Printf("%s %s ", StringReal(left), SEPARATOR)

		// scan the right side of the rule and validate it.
		fmt.Scanf("%s", &right)
		right = functions.Trim(right)

		if reformat(right) != nil {
			right = *reformat(right)
		} else {
			functions.ClearLine(2)
			ErrorPos(right, len(right), "Expected character after "+ESCAPE_CHARACTER)
			continue
		}

		isRValid = ValidateInput(&right, term, nonterm)

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
	fmt.Printf("%s %s %s\n", StringReal(left), SEPARATOR, StringReal(right))
	c.Print("-------------------\n\n")
	// ____________________________________________

	// return the rule as a pointer to the rule.

	return &types.Rule{
		Left:  functions.Split(left),
		Right: functions.Split(right),
	}
}

func reformat(str string) *string {
	var escaped bool = false
	var result string = ""

	for i := 0; i < len(str); i++ {
		char := string((str)[i])

		if char == ESCAPE_CHARACTER && !escaped {
			escaped = true
			continue
		}

		if !escaped {
			if char == NT_SUFFIX {
				result += REAL_NT_SUFFIX
				continue
			}
		} else {
			escaped = false

			if char == NT_SUFFIX {
				result += NT_SUFFIX
				continue
			}
		}

		result += char
	}

	if escaped {
		return nil
	}

	return &result
}

// ValidateInput checks if each character in the input string is present in either
// the term or nonterm slices. If any character is not found in either slice, it
// clears the Rule prompt and returns false, indicating invalid input. Otherwise,
// it returns true.
func ValidateInput(input *string, term []string, nonterm []string) bool {
	var isNonTerm bool = false

	for i := 0; i < len(*input); i++ {
		char := string((*input)[i])

		if char == REAL_NT_SUFFIX {
			// error, unexpected symbol
			functions.ClearLine(2)
			ErrorPos(*input, i, "Unexpected '"+NT_SUFFIX+"'")
			return false
		}

		if i < len(*input)-1 && string((*input)[i+1]) == REAL_NT_SUFFIX {
			isNonTerm = true
			i++
		} else {
			isNonTerm = false
		}

		if isNonTerm && !functions.InArray(nonterm, char) {

			functions.ClearLine(2)
			ErrorPos(*input, i-1, "Symbol not found in non-terminals")
			return false

		} else if !isNonTerm && !functions.InArray(term, char) {

			if functions.InArray(nonterm, char) {
				// inject the suffix to the current position
				*input = functions.InsertAt(*input, REAL_NT_SUFFIX, i+1)
				i--
				continue
			}

			functions.ClearLine(2)
			ErrorPos(*input, i, "Symbol not found in terminals")
			return false
		}

	}

	return true
}

func ErrorPos(str string, pos int, message string) {

	position := pos - 1
	str = StringReal(str)

	if pos == 0 {
		position = 0
	}

	colors.Red.Printf("%s\n", str)
	colors.Red.Printf("%s^\n", strings.Repeat(" ", pos))
	colors.Gray.Printf("%s%s\n", strings.Repeat(" ", position), message)

	fmt.Println()
	fmt.Println()
}
