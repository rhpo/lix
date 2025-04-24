package types

import (
	. "thl/constants"
	. "thl/functions"
)

type Rule struct {
	Left  []string
	Right []string
}

func IsRuleOnlyTerminal(R Rule, term []string) bool {
	for i := 0; i < len(R.Right); i++ {
		if R.Right[i] == REAL_NT_SUFFIX {
			return false
		}
	}

	return true
}

func IsRuleLeftRegular(R Rule, nonterm []string, term []string) bool {
	// The left side of the right of the rule can start with one or more non-terminals
	if len(R.Right) == 0 {
		return true
	}

	// Check if the rule starts with one or more non-terminals followed by NT_SUFFIX
	i := 0
	for i < len(R.Right) && InArray(nonterm, R.Right[i]) {
		i++
	}

	if i > 0 && i < len(R.Right) && R.Right[i] == REAL_NT_SUFFIX {
		// Ensure the rest of the rule contains only terminals
		for j := i + 1; j < len(R.Right); j++ {
			if !InArray(term, R.Right[j]) {
				return false
			}
		}

		return true
	}

	return false
}

func swapStar(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == []rune(REAL_NT_SUFFIX)[0] {
			// Swap '*' with the next character
			runes[i], runes[i+1] = runes[i+1], runes[i]
			i++ // Skip next character to avoid double swaps
		}
	}
	return string(runes)
}

func IsRuleRightRegular(R Rule, nonterm []string, term []string) bool {

	rightSide := Reverse(Join(R.Right))
	rightSide = swapStar(rightSide)

	R_Reversed := &Rule{Left: R.Left, Right: Split(rightSide)}
	return IsRuleLeftRegular(*R_Reversed, nonterm, term)
}

func IsRuleTerminalNonterminal(R Rule, nonterm []string, term []string) bool {

	for i := 0; i < len(R.Right); i++ {
		if !InArray(term, R.Right[i]) && !InArray(nonterm, R.Right[i]) {
			if R.Right[i] == REAL_NT_SUFFIX {
				continue
			}
			return false
		}
	}

	return true
}

func LeftSideGamma(R Rule, nonterm []string, term []string) bool {
	foundNT := false
	result := true

	for i := 0; i < len(R.Left); i++ {
		if R.Left[i] == REAL_NT_SUFFIX {
			continue
		}
		if !InArray(term, R.Left[i]) && !InArray(nonterm, R.Left[i]) {
			result = false
			break
		} else if InArray(nonterm, R.Left[i]) {
			foundNT = true
		}
	}

	return foundNT && result
}
