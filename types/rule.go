package types

import (
	. "thl/functions"
)

const NT_SUFFIX = "*"

type Rule struct {
	Left  []string
	Right []string
}

func IsRuleOnlyTerminal(R Rule, term []string) bool {
	for i := 0; i < len(R.Right); i++ {
		if R.Right[i] == NT_SUFFIX {
			return false
		}
	}

	return true
}

func IsRuleLeftRegular(R Rule, nonterm []string, term []string) bool {
	// the left side of the right of the rule ends with one single non terminal
	if len(R.Right) == 0 {
		return false
	}

	if len(R.Right) == 2 && InArray(nonterm, R.Right[0]) && R.Right[1] == NT_SUFFIX {
		return true
	} else {

		isNonTerminal := len(R.Right) >= 2 && InArray(nonterm, R.Right[0]) && R.Right[1] == NT_SUFFIX

		if isNonTerminal {
			for i := 2; i < len(R.Right); i++ {
				if !InArray(term, R.Right[i]) {
					return false
				}
			}
			return true
		} else {
			return false
		}

	}

}

func swapStar(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if runes[i] == '*' {
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
			if R.Right[i] == NT_SUFFIX {
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
		if R.Left[i] == NT_SUFFIX {
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
