package types

import (
	. "thl/functions"
)

type Rule struct {
	Left  []string
	Right []string
}

func IsRuleOnlyTerminal(R Rule, term []string) bool {
	for i := 0; i < len(R.Right); i++ {
		if !InArray(term, R.Right[i]) {
			return false
		}
	}

	return true
}

func IsRuleLeftRegular(R Rule, nonterm []string, term []string) bool {
	// the left side of the right of the rule ends with one single non terminal
	if len(R.Right) == 0 || len(R.Right) == 1 && InArray(nonterm, R.Right[0]) {
		return true
	} else {
		if InArray(nonterm, R.Right[0]) {
			for i := 1; i < len(R.Right); i++ {
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

func IsRuleRightRegular(R Rule, nonterm []string, term []string) bool {

	R_reversed := Rule{
		Left:  R.Left,
		Right: Split(Reverse(Join(R.Right))),
	}

	return IsRuleLeftRegular(R_reversed, nonterm, term)

}

func IsRuleTerminalNonterminal(R Rule, nonterm []string, term []string) bool {

	for i := 0; i < len(R.Right); i++ {
		if !InArray(term, R.Right[i]) && !InArray(nonterm, R.Right[i]) {
			return false
		}
	}

	return true
}

func LeftSideGamma(R Rule, nonterm []string, term []string) bool {
	foundNT := false
	result := true

	for i := 0; i < len(R.Left); i++ {
		if !InArray(term, R.Left[i]) && !InArray(nonterm, R.Left[i]) {
			result = false
			break
		} else if InArray(nonterm, R.Left[i]) {
			foundNT = true
		}
	}

	return foundNT && result
}
