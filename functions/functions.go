package functions

import "thl/colors"

const helpMessage string = `
__Help_______________________________________
|                                           |
| First,  you will be  prompted  to  enter  |
| the terminals,  non-terminals,  and  the  |
| axiom (start symbol in non-terminals, S). |
|                                           |
| When you finish entering a list of        |
| terminals or non-terminals, press Enter.  |
|                                           |
| Then, you will Enter the Rules...         |
| Non-terminals are followed by '*'         |
|                                           |
| When you finish entering the rules, press |
| Enter to generate the language.           |
|                                           |
| where  A and B  are  strings of terminals |
| and non-terminals.                        |
|                                           |
‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾`

func Help() {
	Clear()

	colors.Blue.Print("\nWelcome to the Language Generator\n")
	colors.Gray.Println(helpMessage)
}
