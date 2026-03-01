// ============================================================================
// HELIX EDITOR TUTOR — Learn by Editing Go Code
// ============================================================================
//
// Open this file with: hx ~/helix-tutor/tutor.go
//
// HOW TO USE:
// - Each section teaches one concept
// - Read the TASK, then do it using the COMMAND shown
// - The code has intentional bugs/gaps for you to fix
// - Work top to bottom, save after each section with :w or Ctrl-s
//
// TIP: Press Space in normal mode to see all available commands!
//      Press g to see the goto menu!
//      These menus are your best friends as a beginner.
//
// YOUR CUSTOM BINDINGS (safe, no omarchy conflicts):
//   Ctrl-s     = save (works in normal AND insert mode)
//   Ctrl-q     = close current buffer
//   Ctrl-d/u   = half page down/up (auto-centers screen)
//   jk         = escape to normal mode (from insert)
//   U          = redo (opposite of u for undo)
//   X          = extend line selection upward
//   Space w    = save        Space q   = close buffer
//   Space g s  = lazygit     Space g d = git diff
//   Space e    = next error  Space E   = prev error
//   Space y    = copy to system clipboard
//   Space p    = paste from system clipboard
//   Space i    = toggle inlay hints
//   Space n    = toggle relative/absolute line numbers
//   Space t    = open terminal (tmux split below)
//   Space l r  = restart LSP
//   Space l f  = format file
//   Space l a  = code action
//   Space l R  = rename symbol
//   Space l d  = goto definition
//   Space l D  = goto declaration
//   Space l i  = goto implementation
//   Space l t  = goto type definition
//   Space l s  = signature help
//   Space l h  = hover docs
//
// ============================================================================

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ============================================================================
// LESSON 1: BASIC MOVEMENT
// ============================================================================
//
// COMMANDS:
//   h = left    j = down    k = up    l = right
//   w = next word start     b = previous word start
//   e = next word end
//   gh = line start         gl = line end
//   gg = file start         ge = file end
//
// TASK: Navigate to each comment below and read them in order.
// Just practice moving around. No editing yet.

func lesson1Movement() {
	first := "start here"     // 1. You are here
	second := "move down"     // 2. Press j to come here
	third := "go right"       // 3. Press l to move right along this line -------->
	fourth := "word jump"     // 4. Press w w w to jump word by word
	fifth := "go back"        // 5. Press b b b to go back word by word
	_ = first
	_ = second
	_ = third
	_ = fourth
	_ = fifth
}

// ============================================================================
// LESSON 2: VERTICAL MOVEMENT & JUMPING
// ============================================================================
//
// COMMANDS:
//   Ctrl-d = half page down (auto-centers!)
//   Ctrl-u = half page up (auto-centers!)
//   5j     = jump 5 lines down (use relative line numbers!)
//   12k    = jump 12 lines up
//   :<num> = go to line number (e.g. :100)
//
// WHY RELATIVE NUMBERS: Look at the left gutter. The numbers show
// DISTANCE from your cursor. To jump to a line, just type the number + j/k.
//
// TASK: Use Ctrl-d to scroll down to lesson 3, then Ctrl-u back here.
// Then use the relative numbers to jump: e.g. 8j to go 8 lines down.

func lesson2Jumping() {
	fmt.Println("Use Ctrl-d to page down")
	fmt.Println("Use Ctrl-u to page up")
	fmt.Println("Both auto-center your screen!")
	fmt.Println("---")
	fmt.Println("Now look at relative line numbers on the left")
	fmt.Println("If this line shows '3' on the left, type 3j to jump here")
	fmt.Println("---")
	fmt.Println("Practice: jump between these lines using number + j/k")
}

// ============================================================================
// LESSON 3: ENTERING INSERT MODE
// ============================================================================
//
// COMMANDS:
//   i = insert before cursor     a = insert after cursor
//   I = insert at line start     A = insert at line end
//   o = new line below           O = new line above
//   Escape or jk = back to normal mode
//
// TASK: Fix the function below. Each variable needs a value.
// TIP: Use 'jk' instead of Escape — faster!

func lesson3Insert() {
	name :=    // FIX: press A to go to end, add "gopher"
	age :=     // FIX: press A, add 10
	language := // FIX: press A, add "Go"

	fmt.Println(name, age, language)
}

// ============================================================================
// LESSON 4: SELECTION — The Helix Superpower
// ============================================================================
//
// THIS IS THE BIG DIFFERENCE FROM VIM!
// In Helix, you SELECT first, then ACT. (Vim does verb-object)
//
// Vim:   d w     = delete word  (verb first)
// Helix: w d     = select word, then delete (selection first)
//
// You always SEE what you're about to affect!
//
// COMMANDS:
//   v = enter select mode (extend selection with movement)
//   x = select entire line
//   X = extend selection one line UP (custom binding)
//   ; = collapse selection to cursor
//   s = search within selection (super powerful!)
//   % = select entire file
//
// TASK: Select each line with 'x', then 'd' to delete the REMOVE lines.

func lesson4Selection() {
	keepThis := "important"
	REMOVE_THIS_LINE := "trash"
	alsoKeep := "valuable"
	REMOVE_THIS_LINE_TOO := "garbage"
	keepThisToo := "essential"
	REMOVE_ME_ALSO := "junk"

	_ = keepThis
	_ = REMOVE_THIS_LINE
	_ = alsoKeep
	_ = REMOVE_THIS_LINE_TOO
	_ = keepThisToo
	_ = REMOVE_ME_ALSO
}

// ============================================================================
// LESSON 5: SELECTION + ACTION (The Helix Way)
// ============================================================================
//
// COMMANDS:
//   w  = select next word          then d = delete it
//   w  = select next word          then c = change it (delete + insert)
//   x  = select line               then d = delete line
//   x  = select line               then y = yank (copy)
//   p  = paste after               P = paste before
//
// REMEMBER: motion = selection. Then you act on it.
//
// TASK: Fix the function using selection + action

func lesson5SelectionAction() {
	animal := "WRONG"   // use w on WRONG, then c, type "cat"
	color := "WRONG"    // use w on WRONG, then c, type "blue"
	number := 0         // use w on 0, then c, type 42

	fmt.Println(animal, color, number)
}

// ============================================================================
// LESSON 6: FIND AND SEARCH
// ============================================================================
//
// COMMANDS:
//   f<char> = find character forward (select up to and including it)
//   t<char> = find character forward (select up to but NOT including it)
//   F/T     = same but backward
//   /       = search forward in file
//   n       = next result (auto-centers with your config!)
//   N       = previous result (auto-centers!)
//   *       = search for word under cursor
//
// TASK: Use f and / to navigate and fix the bugs

func lesson6FindSearch() {
	msg1 := "Helo, Wrld!"         // should be "Hello, World!"
	msg2 := "Go is awsome!"       // should be "Go is awesome!"
	msg3 := "Helix is amazng!"    // should be "Helix is amazing!"

	// Use / to search for "BUG" then n to jump between them
	result1 := 10 + BUG   // BUG should be 5
	result2 := 20 * BUG   // BUG should be 3
	result3 := BUG - 1    // BUG should be 100

	fmt.Println(msg1, msg2, msg3, result1, result2, result3)
}

// ============================================================================
// LESSON 7: MULTIPLE CURSORS (Helix Superpower!)
// ============================================================================
//
// COMMANDS:
//   C         = add cursor on line below
//   Alt-C     = add cursor on line above
//   s         = select all regex matches within selection
//   &         = align cursors
//   ,         = keep only primary selection
//   Alt-,     = remove primary selection
//   (         = rotate selections backward
//   )         = rotate selections forward
//
// TASK 1: Select all 5 lines below with 5x, then press 's', type 'foo',
//         Enter, then 'c' to change all 'foo' to 'bar' at once!

func lesson7MultiCursor() {
	foo1 := "foo"
	foo2 := "foo"
	foo3 := "foo"
	foo4 := "foo"
	foo5 := "foo"

	// TASK 2: Place cursor on first 'old', press C C C to add cursors below
	// then use 'w' to select 'old', then 'c' to type 'new'
	old_value1 := 1
	old_value2 := 2
	old_value3 := 3
	old_value4 := 4

	_ = foo1; _ = foo2; _ = foo3; _ = foo4; _ = foo5
	_ = old_value1; _ = old_value2; _ = old_value3; _ = old_value4
}

// ============================================================================
// LESSON 8: SPACE MENU — Your Command Center
// ============================================================================
//
// Press Space in normal mode to see ALL available commands!
//
// DEFAULT SPACE COMMANDS:
//   Space f   = file picker (fuzzy find files!)
//   Space b   = buffer picker (switch open files)
//   Space k   = show docs for symbol under cursor
//   Space r   = rename symbol (LSP)
//   Space a   = code action (LSP)
//   Space d   = diagnostics picker
//   Space /   = global search
//   Space s   = symbol picker
//   Space S   = workspace symbol picker
//   Space h   = select references (LSP)
//
// YOUR CUSTOM SPACE COMMANDS:
//   Space w   = save file
//   Space q   = close buffer
//   Space e   = next diagnostic    Space E = prev diagnostic
//   Space y   = yank to clipboard  Space p = paste from clipboard
//   Space i   = toggle inlay hints
//   Space n   = toggle line number style
//   Space t   = open terminal below (tmux split)
//   Space g s = lazygit
//   Space g d = git diff
//   Space g l = git log
//   Space g b = git blame current file
//   Space l r = restart LSP
//   Space l f = format file
//   Space l a = code action
//   Space l R = rename symbol
//   Space l d/D = goto definition/declaration
//   Space l i = goto implementation
//   Space l t = goto type definition
//   Space l s = signature help
//   Space l h = hover docs
//
// TASK: Try Space then look at the menu. Press Escape to cancel.
// Then try Space f to browse files.

func lesson8SpaceMenu() {
	myVariable := "try Space k on fmt.Println to see docs"
	fmt.Println(myVariable)
	// Try Space r on myVariable to rename it everywhere
	// Try Space e to jump to the next error in the file
}

// ============================================================================
// LESSON 9: g MENU AND GOTO
// ============================================================================
//
// COMMANDS:
//   g         = goto menu (press g then see options)
//   gd        = go to definition
//   gy        = go to type definition
//   gr        = go to references
//   gi        = go to implementation
//   ga        = go to last accessed file (toggle between 2 files!)
//   gn        = go to next buffer
//   gp        = go to previous buffer
//   gl        = go to end of line
//   gh        = go to start of line
//   gs        = go to first non-blank in line
//
//   Ctrl-o    = jump back (after any goto)
//   Ctrl-i    = jump forward
//
// TASK: Use gd on Println to jump to its definition, then Ctrl-o back

func lesson9Goto() {
	fmt.Println("Use gd on Println to jump to definition")
	fmt.Println("Then Ctrl-o to jump back")
	fmt.Println("Try ga to toggle between this and the last file you visited")
}

// ============================================================================
// LESSON 10: TEXT OBJECTS (Power Moves)
// ============================================================================
//
// Helix uses mi/ma for "match inside" / "match around"
//
//   miw = select inner word       maw = select around word (+ space)
//   mi" = select inside quotes    ma" = select around quotes
//   mi( = select inside parens    ma( = select around parens
//   mi{ = select inside braces    ma{ = select around braces
//   mif = select inside function  maf = select around function
//
// THEN act on it: d (delete), c (change), y (yank)
//
// TASK: Use text objects to edit efficiently

func lesson10TextObjects() {
	// Use mi" then c to change the string content
	greeting := "change this text to hello world"

	// Use mi( then d to delete the arguments
	fmt.Println("delete", "these", "args")

	// Use maf to select this entire function, then y to yank it
	fmt.Println(greeting)
}

// ============================================================================
// LESSON 11: MATCH MODE AND SURROUND
// ============================================================================
//
// COMMANDS:
//   mm        = jump to matching bracket
//   ms<char>  = surround selection with char
//   md<char>  = delete surrounding char
//   mr<old><new> = replace surrounding char
//
// TASK: Fix the brackets and surrounds

func lesson11MatchSurround() {
	// Use mm to jump between matching brackets
	data := map[string][]int{
		"scores": {90, 85, 92, 88, 95},
		"ages":   {25, 30, 35, 40, 45},
	}

	// Select word, then ms" to surround with quotes
	bare := hello    // should be "hello"

	// Select content including quotes, then md" to remove them
	extra := ""world""  // should be just: world

	fmt.Println(data, bare, extra)
}

// ============================================================================
// LESSON 12: WINDOW SPLITS (Inside Helix)
// ============================================================================
//
// COMMANDS:
//   Ctrl-w s  = horizontal split
//   Ctrl-w v  = vertical split
//   Ctrl-w q  = close split
//   Ctrl-w w  = cycle to next split
//   Ctrl-w h/j/k/l = move to split in direction
//
// NO CONFLICTS:
//   Hyprland uses Super+arrows for window focus
//   Tmux uses Ctrl-Alt-arrows for pane focus
//   Helix uses Ctrl-w for internal splits
//   All three layers, zero overlap!
//
// TASK: Ctrl-w v to split, Space f to open a file, Ctrl-w q to close

func lesson12Splits() {
	fmt.Println("Split this view with Ctrl-w v")
	fmt.Println("Open another file with Space f in the new split")
	fmt.Println("Close split with Ctrl-w q")
}

// ============================================================================
// LESSON 13: REGISTERS AND MACROS
// ============================================================================
//
// COMMANDS:
//   "xy       = yank into register x
//   "xp       = paste from register x
//   Q         = start/stop recording macro
//   q         = play recorded macro
//
// SYSTEM CLIPBOARD (your custom bindings):
//   Space y   = yank to system clipboard
//   Space p   = paste from system clipboard
//
// TASK: Record a macro to duplicate a line
// 1. Go to "item 1", press Q to start recording
// 2. x (select line), y (yank), p (paste below)
// 3. Press Q to stop
// 4. Press q to replay and create more items

func lesson13Registers() {
	items := []string{
		"item 1",
		// Use macro to duplicate and create item 2, 3, 4, 5
	}
	fmt.Println(items)
}

// ============================================================================
// LESSON 14: LSP FEATURES WITH Go
// ============================================================================
//
// With gopls running, you get:
//
//   Space k   = hover docs
//   Space r   = rename symbol
//   Space a   = code actions
//   gd        = go to definition
//   gr        = go to references
//   Space d   = diagnostics picker
//   Space e   = next diagnostic (custom)
//   Space E   = prev diagnostic (custom)
//   Space l r = restart LSP (custom)
//   Space l f = format file (custom)
//   Space l a = code action (custom)
//   Space l R = rename symbol (custom)
//   Space l d = goto definition (custom)
//   Space l D = goto declaration (custom)
//   Space l i = goto implementation (custom)
//   Space l t = goto type definition (custom)
//   Space l s = signature help (custom)
//   Space l h = hover docs (custom)
//
// Inlay hints show parameter names and types inline!
// Toggle them with Space i (custom)
//
// TASK: Use LSP features to explore and fix this code

func lesson14LSP() {
	// Hover over Sprintf with Space k to see its signature
	msg := fmt.Sprintf("Hello %s, you are %d years old", "Gopher", 5)

	// Use Space e to jump to the next error
	// Then Space a to see code actions
	result := fmt.Sprintff("broken")  // typo! fix it

	// Rename 'oldName' to 'userName' using Space r
	oldName := "Gopher"
	fmt.Println(msg, result, oldName)
}

// ============================================================================
// LESSON 15: COMMAND MODE
// ============================================================================
//
// Press : to enter command mode
//
// USEFUL COMMANDS:
//   :w          = save (or Ctrl-s, or Space w)
//   :q          = quit
//   :wq         = save and quit
//   :q!         = force quit (or Space Q)
//   :bc         = close buffer (or Ctrl-q, or Space q)
//   :bn / :bp   = next/prev buffer (or gn/gp)
//   :theme      = change theme (Tab to see options)
//   :set        = change settings live
//   :sh <cmd>   = run shell command
//   :config-reload = reload config without restarting
//   :tutor      = built-in helix tutor
//
// TASK: Try :theme and Tab through options. Try :sh go version

func lesson15Commands() {
	fmt.Println("Try :theme catppuccin_mocha")
	fmt.Println("Try :sh go version")
	fmt.Println("Try :config-reload after editing config")
}

// ============================================================================
// LESSON 16: GIT INTEGRATION (Custom Bindings)
// ============================================================================
//
// Your config has git shortcuts under Space g:
//   Space g s = open lazygit (full TUI git client)
//   Space g d = git diff
//   Space g l = git log (last 20 commits)
//   Space g b = git blame current file
//
// The gutter also shows diff markers:
//   green  = added lines
//   yellow = modified lines
//   red    = deleted lines
//
// TASK: If this is a git repo, try Space g s to open lazygit
// Navigate lazygit: q to quit, ? for help

func lesson16Git() {
	fmt.Println("Space g s opens lazygit")
	fmt.Println("Space g d shows diff")
	fmt.Println("Space g l shows log")
}

// ============================================================================
// LESSON 17: TMUX + HELIX WORKFLOW
// ============================================================================
//
// Your terminal has 3 layers, all with different keys:
//
//   Hyprland:  Super + keys          (desktop windows)
//   Tmux:      Ctrl-Space + keys     (terminal multiplexer)
//   Helix:     Space / g / : / Ctrl-w (editor)
//
// QUICK WORKFLOW:
//   Space t    = open terminal below (tmux pane)
//   Then run: go test ./... or go run .
//   Ctrl-Alt-Up = go back to helix pane
//
// FULL WORKFLOW:
//   t              = start tmux
//   Ctrl-Space v   = split vertical
//   Left pane:     hx .
//   Right pane:    go test, go run, git
//   Alt-1..9       = switch tmux windows
//   Alt-Up/Down    = switch tmux sessions
//
// TASK: Try Space t to open a terminal, run 'go version', then
// Ctrl-Alt-Up to come back to helix

func lesson17TmuxWorkflow() {
	fmt.Println("Press Space t to open a terminal below")
	fmt.Println("Run: go version")
	fmt.Println("Press Ctrl-Alt-Up to come back")
}

// ============================================================================
// LESSON 18: PIPE THROUGH SHELL (Advanced)
// ============================================================================
//
// Select text, then pipe it through a shell command:
//   | <cmd>  = pipe selection through command, replace with output
//   ! <cmd>  = pipe selection, insert output (keep original)
//
// TASK: Select the unsorted lines below (use x on each or v to select),
// then type: | sort

func lesson18Pipe() {
	// Select these lines and pipe through sort
	_ = []string{
		"delta",
		"alpha",
		"charlie",
		"echo",
		"bravo",
	}
}

// ============================================================================
// FINAL BOSS: Debug a Go Program
// ============================================================================
//
// Use everything you learned to fix this program.
// It should print a sorted list of unique words from the input.
// Use Space e to jump between errors, gd to explore, Space a for fixes.

func finalBoss() {
	input := "the quick brown fox jumps over the lazy dog the fox"

	// BUG 1: wrong function name
	words := strings.Splitt(input, " ")

	// BUG 2: wrong type for map
	seen := make(map[int]bool)

	// BUG 3: logic error in deduplication
	unique := []string{}
	for _, w := range words {
		if seen[w] {  // should check if NOT seen
			unique = append(unique, w)
			seen[w] = true
		}
	}

	// BUG 4: wrong sort function
	sort.string(unique)

	fmt.Println("Unique words:", unique)
}

// ============================================================================
// CHEAT SHEET
// ============================================================================
//
// ╔═══════════════════════════════════════════════════════════════╗
// ║  THE HELIX WAY: Select first → then Act                     ║
// ╠═══════════════════════════════════════════════════════════════╣
// ║  MOVEMENT     h j k l   w b e   gg ge   f/t   /  Ctrl-d/u  ║
// ║  SELECTION    v x X %   s (search in sel)   ; (collapse)    ║
// ║  ACTIONS      d c y p   > < (indent)   ~ (case)   U (redo) ║
// ║  MULTICURSOR  C (below)  s (matches)  , (keep primary)     ║
// ║  TEXT OBJ     mi/ma + w " ( { [   mif/maf (function)       ║
// ║  SURROUND     ms (add)  md (delete)  mr (replace)          ║
// ║  GOTO         g → gd gr gi gl gh   Ctrl-o/i (jump)        ║
// ║  SPACE        f b k r a d / s   w q e y p i n t g l       ║
// ║  SPLITS       Ctrl-w → s v q w h/j/k/l                    ║
// ║  COMMANDS     : → w q wq bc theme sh config-reload         ║
// ║  SAVE         Ctrl-s (any mode)   Space w   :w             ║
// ║  ESCAPE       Escape or jk (from insert)                   ║
// ╚═══════════════════════════════════════════════════════════════╝
//
// ============================================================================

func main() {
	lesson1Movement()
	lesson2Jumping()
	lesson3Insert()
	lesson4Selection()
	lesson5SelectionAction()
	lesson6FindSearch()
	lesson7MultiCursor()
	lesson8SpaceMenu()
	lesson9Goto()
	lesson10TextObjects()
	lesson11MatchSurround()
	lesson12Splits()
	lesson13Registers()
	lesson14LSP()
	lesson15Commands()
	lesson16Git()
	lesson17TmuxWorkflow()
	lesson18Pipe()
	finalBoss()
}
