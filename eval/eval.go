package eval

import (
	. "thl/constants"
	. "thl/functions"
	. "thl/types"
)

const MAX_RECURSION = 5000

type Slot struct {
	// The slot is a string that can be a terminal or non-terminal
	super bool
	value string
}

type SuperRule struct {
	left  []Slot
	right []Slot
}

func translateSide(g Grammar, arr []string) []Slot {
	side := make([]Slot, 0)

	i := 0
	for i < len(arr) {
		char := arr[i]

		if i+1 < len(arr) && arr[i+1] == REAL_NT_SUFFIX && InArray(g.N, char) {
			// Current char + suffix → it's a non-terminal
			side = append(side, Slot{
				super: true,
				value: char,
			})
			i += 2 // skip the suffix
		} else if InArray(g.T, char) {
			// Terminal
			side = append(side, Slot{
				super: false,
				value: char,
			})
			i++
		} else {
			// Unknown char, just skip
			i++
		}
	}

	return side
}

func translator(g Grammar) []SuperRule {
	var superRules []SuperRule = make([]SuperRule, 0)

	// Supposing that all the grammar rules are valid from the previous step.

	for _, rule := range g.P {

		left := translateSide(g, rule.Left)
		right := translateSide(g, rule.Right)

		superRules = append(superRules, SuperRule{
			left:  left,
			right: right,
		})
	}

	return superRules
}

func findRulesByNonTerminal(rules []SuperRule, nt string) []SuperRule {
	var result []SuperRule
	for _, rule := range rules {
		if len(rule.left) == 1 && rule.left[0].super && rule.left[0].value == nt {
			result = append(result, rule)
		}
	}
	return result
}

func evaluate(rules []SuperRule, current string, min, max, depth int) []string {
	if depth > MAX_RECURSION {
		return []string{}
	}

	if len(current) > max {
		return []string{}
	}

	var results []string

	hasNonTerminal := false
	for i := 0; i < len(current); i++ {
		c := string(current[i])
		if isNonTerminal(c, rules) {
			hasNonTerminal = true
			for _, rule := range findRulesByNonTerminal(rules, c) {
				newString := current[:i] + slotsToString(rule.right) + current[i+1:]
				results = append(results, evaluate(rules, newString, min, max, depth+1)...)
			}
			break
		}
	}

	if !hasNonTerminal && len(current) >= min {
		results = append(results, current)
	}

	return results
}

func isNonTerminal(c string, rules []SuperRule) bool {
	for _, rule := range rules {
		if len(rule.left) > 0 && rule.left[0].value == c && rule.left[0].super {
			return true
		}
	}
	return false
}

func Evaluate(g Grammar, min, max int) []string {
	superRules := translator(g)
	startSymbol := g.S
	return evaluate(superRules, startSymbol, min, max, 0)
}

// Helper function to convert a slice of Slots to a string
func slotsToString(slots []Slot) string {
	var result string
	for _, slot := range slots {
		result += slot.value
	}
	return result
}
