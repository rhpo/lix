
# lix 📝⚡
A **language grammar generator** inspired by Chomsky's formal grammars.

## 🚀 Features
- Define and manipulate **context-free grammars (CFGs)**
- Generate **syntactically valid sentences**
- Analyze and transform grammatical structures
- Simple CLI for quick input and generation

## 📖 Usage

### 1️⃣ Enter Grammar Components
You'll be prompted to enter:
- **Terminals** (symbols in the language)
- **Non-terminals** (variables in the grammar)
- **Axiom** (start symbol, usually `S`)

Press **Enter** after each list.

### 2️⃣ Define Production Rules
- Each **non-terminal** is followed by `*`.
- Rules must map **non-terminals** to **terminals/non-terminals**.
- When done, press **Enter** to generate the language.

### Example Input
```
Enter terminals: a b c
Enter non-terminals: S A B
Enter axiom: S

Enter rules:
S → A B
A* → a A | a
B* → b B | b c
```

### 3️⃣ Generate & Explore
Once rules are entered, `lix` will generate and analyze valid sentences from your grammar.

## 🔗 Coming Soon
- Documentation & examples
- CLI enhancements
