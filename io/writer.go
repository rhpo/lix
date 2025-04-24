package io

import (
	"fmt"
	"strings"

	. "thl/constants"

	"thl/colors"
	"thl/functions"
	"thl/types"
)

func PrintTable(A []string, title string, color types.IColor, axioms ...string) {

	var axiom string = ""
	if len(axioms) > 0 {
		axiom = axioms[0]
	}

	A = functions.Sort(A)

	// if A has the axiom, remove it from the list and add it to the start
	if functions.InArray(A, axiom) {
		A = functions.Remove(A, axiom)
		A = append([]string{axiom}, A...)
	}

	var header string = title
	color.Println(header)

	color.Print(STR_WALL)
	for i := 0; i < len(A); i++ {
		A[i] = StringReal(A[i])

		if i == len(A)-1 {
			color.Printf("%s\n", A[i])
		} else if A[i] == axiom {
			colors.CyanBG.Printf(" %s ", A[i])
			fmt.Print(" ")
		} else {
			color.Printf("%s, ", A[i])
		}
	}

	color.Println(strings.Repeat("=", len(header)+1))
	fmt.Println()
}

func PrintRules(R []types.Rule, nonterm []string, color types.IColor) {
	const header string = "Rules"
	var longestRule int = len(header)
	var ruleLength int = 0

	color.Println(header)

	for i := 0; i < len(R); i++ {
		color.Print(STR_WALL)
		ruleLength = 0 // Reset ruleLength for each rule

		for j := 0; j < len(R[i].Left); j++ {
			l := R[i].Left[j]

			// if next character is equal to REAL_NT_SUFFIX, colorize the current character
			if j < len(R[i].Left)-1 && R[i].Left[j+1] == REAL_NT_SUFFIX {

				colors.Cyan.Printf("%s"+NT_SUFFIX, l)

				j++ // for the axiom
				j++ // for the NT_SUFFIX
				ruleLength++
			} else {
				color.Printf("%s", l)
			}

			ruleLength++
		}

		color.Print(" ", SEPARATOR, " ")
		ruleLength += len(SEPARATOR) + 2

		if functions.Join(R[i].Right) == "" {
			colors.Gray.Println("ε")
		} else {
			for j := 0; j < len(R[i].Right); j++ {
				r := R[i].Right[j]

				if j < len(R[i].Right)-1 && R[i].Right[j+1] == REAL_NT_SUFFIX {
					colors.Cyan.Printf("%s", r)

					j++
					ruleLength++
				} else {
					color.Printf("%s", r)
				}

				ruleLength++
			}

			fmt.Println()
		}

		if ruleLength > longestRule {
			longestRule = ruleLength
		}
	}

	color.Println(strings.Repeat("=", longestRule+1))
}
