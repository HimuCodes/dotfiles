# Tmux + Helix on Omarchy — The Complete Guide

## The Big Picture

Your workflow has three layers. Each uses different keys. Zero conflicts.

```
┌─────────────────────────────────────────────────┐
│  HYPRLAND (Omarchy)     — Super + keys          │
│  ┌───────────────────────────────────────────┐   │
│  │  TMUX                — Ctrl-Space + keys  │   │
│  │  ┌─────────────────────────────────────┐  │   │
│  │  │  HELIX            — Space / g / :   │  │   │
│  │  └─────────────────────────────────────┘  │   │
│  └───────────────────────────────────────────┘   │
└─────────────────────────────────────────────────┘
```

---

## Layer 1: Hyprland (Omarchy) — Desktop

| Key | Action |
|-----|--------|
| `Super + Return` | New terminal |
| `Super + Alt + Return` | New terminal in tmux |
| `Super + W` | Close window |
| `Super + arrows` | Move window focus |
| `Super + 1-9` | Switch workspace |
| `Super + Shift + N` | Open editor |
| `Super + Space` | App launcher |
| `Super + Alt + Space` | Omarchy menu |
| `Super + F` | Fullscreen |
| `Super + T` | Toggle floating |
| `Super + Tab / Shift+Tab` | Next/prev workspace |
| `Super + S` | Scratchpad |
| `Super + L` | Toggle scrolling layout |
| `Super + K` | Show keybindings |
| `Super + G` | Toggle window grouping |
| `Super + J` | Toggle split |
| `Super + O` | Pop window out (float + pin) |
| `Super + C` | Copy (universal) |
| `Super + V` | Paste (universal) |
| `Super + Ctrl + V` | Clipboard manager |
| `Super + Comma` | Dismiss notification |
| `Super + Escape` | System menu |
| `Super + /` | Cycle display resolution |
| `Super + Ctrl + O` | Toggle menu (gaps, ratio, scaling) |
| `Super + Ctrl + L` | Lock screen |
| `Super + Ctrl + A/B/W` | Audio / Bluetooth / WiFi controls |
| `PrintScr` | Screenshot |
| `Alt + PrintScr` | Screen recording |

**Rule:** Super key = Hyprland. Helix never sees Super. No conflicts possible.

---

## Layer 2: Tmux — Terminal Multiplexer

### Starting Tmux

```bash
t                              # attach or new session (omarchy alias)
Super + Alt + Return           # terminal directly in tmux
tmux new -s project            # named session
tdl cx                         # dev layout: editor + claude + terminal
tdlm cx                        # dev layout per subdirectory
tsl 4 cx                       # swarm: 4 panes all running claude
```

### Prefix: `Ctrl-Space`

Press `Ctrl-Space`, release, then the command key.

### Panes

| Key | Action |
|-----|--------|
| `Ctrl-Space h` | Split horizontal |
| `Ctrl-Space v` | Split vertical |
| `Ctrl-Space x` | Close pane |
| `Ctrl-Alt-arrows` | Move between panes |
| `Ctrl-Alt-Shift-arrows` | Resize panes |

### Windows

| Key | Action |
|-----|--------|
| `Ctrl-Space c` | New window |
| `Ctrl-Space k` | Kill window |
| `Ctrl-Space r` | Rename window |
| `Alt-1` through `Alt-9` | Jump to window |
| `Alt-Left/Right` | Prev/next window |

### Sessions

| Key | Action |
|-----|--------|
| `Ctrl-Space C` | New session |
| `Ctrl-Space K` | Kill session |
| `Ctrl-Space R` | Rename session |
| `Alt-Up/Down` | Switch sessions |
| `Ctrl-Space d` | Detach (session keeps running!) |
| `tmux attach` | Reattach later |

### Copy Mode

| Key | Action |
|-----|--------|
| `Ctrl-Space [` | Enter copy mode |
| `v` | Start selection |
| `y` | Copy |
| `Ctrl-Space ]` | Paste |

---

## Layer 3: Helix — Your Editor

### Opening Files

```bash
hx file.go                     # single file
hx .                           # file picker for current dir
hx src/                        # file picker for directory
hx file1.go file2.go           # multiple buffers
```

### The Helix Philosophy

**Select → Act** (opposite of Vim)

```
Vim:   dw      = delete word  (verb → noun)
Helix: wd      = select word, then delete  (see it → do it)
```

You always SEE what you'll affect before acting.

### Your Custom Keybindings

These are verified conflict-free with Omarchy + Tmux:

| Key | Mode | Action |
|-----|------|--------|
| `Ctrl-s` | Normal + Insert | Save file |
| `Ctrl-q` | Normal | Close buffer |
| `Ctrl-d` | Normal | Half page down + center |
| `Ctrl-u` | Normal | Half page up + center |
| `n / N` | Normal | Next/prev search + center |
| `jk` | Insert | Escape to normal mode |
| `U` | Normal | Redo |
| `X` | Normal + Select | Extend line selection up |
| `Space w` | Normal | Save |
| `Space W` | Normal | Save all |
| `Space q` | Normal | Close buffer |
| `Space Q` | Normal | Force quit |
| `Space e` | Normal | Next diagnostic |
| `Space E` | Normal | Previous diagnostic |
| `Space y` | Normal | Yank to system clipboard |
| `Space p` | Normal | Paste from system clipboard |
| `Space P` | Normal | Paste before from clipboard |
| `Space i` | Normal | Toggle inlay hints |
| `Space n` | Normal | Toggle relative/absolute line numbers |
| `Space t` | Normal | Open terminal below (tmux split) |
| `Space g s` | Normal | Open lazygit |
| `Space g d` | Normal | Git diff |
| `Space g l` | Normal | Git log (20 commits) |
| `Space g b` | Normal | Git blame current file |
| `Space l r` | Normal | Restart LSP |
| `Space l f` | Normal | Format file |
| `Space l a` | Normal | Code action |
| `Space l R` | Normal | Rename symbol |
| `Space l d` | Normal | Go to definition |
| `Space l D` | Normal | Go to declaration |
| `Space l i` | Normal | Go to implementation |
| `Space l t` | Normal | Go to type definition |
| `Space l s` | Normal | Signature help |
| `Space l h` | Normal | Hover docs |

### Default Helix Commands Reference

#### Movement
| Key | Action |
|-----|--------|
| `h j k l` | Left, down, up, right |
| `w / b / e` | Next word / prev word / end word |
| `gh / gl` | Line start / line end |
| `gs` | First non-blank |
| `gg / ge` | File start / file end |
| `f<char> / F<char>` | Find char forward / backward |
| `t<char> / T<char>` | Till char forward / backward |
| `/ / ?` | Search forward / backward |
| `n / N` | Next / prev match (auto-centers) |
| `*` | Search word under cursor |
| `Ctrl-d / Ctrl-u` | Half page down/up (auto-centers) |
| `5j / 12k` | Jump N lines (use relative numbers!) |

#### Selection
| Key | Action |
|-----|--------|
| `v` | Enter select mode |
| `x` | Select line |
| `X` | Extend selection up |
| `%` | Select entire file |
| `s` | Search within selection |
| `;` | Collapse selection to cursor |
| `C` | Add cursor below |
| `Alt-C` | Add cursor above |
| `,` | Keep only primary cursor |

#### Actions
| Key | Action |
|-----|--------|
| `d` | Delete |
| `c` | Change (delete + insert) |
| `y` | Yank (copy) |
| `p / P` | Paste after / before |
| `u / U` | Undo / redo |
| `> / <` | Indent / dedent |
| `~` | Toggle case |
| `J` | Join lines |
| `.` | Repeat last change |

#### Text Objects
| Key | Action |
|-----|--------|
| `miw / maw` | Inside / around word |
| `mi" / ma"` | Inside / around quotes |
| `mi( / ma(` | Inside / around parens |
| `mi{ / ma{` | Inside / around braces |
| `mi[ / ma[` | Inside / around brackets |
| `mif / maf` | Inside / around function |
| `mic / mac` | Inside / around class |

#### Surround
| Key | Action |
|-----|--------|
| `ms<char>` | Add surround |
| `md<char>` | Delete surround |
| `mr<old><new>` | Replace surround |
| `mm` | Jump to matching bracket |

#### Space Menu
| Key | Action |
|-----|--------|
| `Space f` | File picker |
| `Space b` | Buffer picker |
| `Space k` | Hover docs |
| `Space r` | Rename symbol |
| `Space a` | Code actions |
| `Space d` | Diagnostics picker |
| `Space /` | Global search |
| `Space s` | Symbol picker |
| `Space S` | Workspace symbols |
| `Space h` | Select references |

#### Goto Menu (press g)
| Key | Action |
|-----|--------|
| `gd` | Go to definition |
| `gy` | Go to type definition |
| `gr` | Go to references |
| `gi` | Go to implementation |
| `ga` | Last accessed file (toggle!) |
| `gn / gp` | Next / prev buffer |
| `gl / gh` | Line end / start |
| `Ctrl-o / Ctrl-i` | Jump back / forward |

#### Splits
| Key | Action |
|-----|--------|
| `Ctrl-w v` | Vertical split |
| `Ctrl-w s` | Horizontal split |
| `Ctrl-w q` | Close split |
| `Ctrl-w w` | Cycle splits |
| `Ctrl-w h/j/k/l` | Focus split |

#### Command Mode (`:`)
| Command | Action |
|---------|--------|
| `:w` | Save |
| `:q` | Quit |
| `:wq` | Save + quit |
| `:q!` | Force quit |
| `:bc` | Close buffer |
| `:bn / :bp` | Next/prev buffer |
| `:theme <name>` | Change theme |
| `:sh <cmd>` | Run shell command |
| `:config-reload` | Reload config |
| `:tutor` | Built-in tutor |

#### Pipe Through Shell
| Key | Action |
|-----|--------|
| `\| sort` | Sort selected lines |
| `\| uniq` | Remove duplicate lines |
| `\| wc -l` | Count lines |
| `! cmd` | Insert command output |

---

## Recommended Workflows

### Workflow 1: Quick Edit

```bash
hx file.go         # open, edit, Ctrl-s to save, :q to quit
```

### Workflow 2: Solo Dev (Most Common)

```
t                              # start tmux
Ctrl-Space v                   # split vertical
# Left:  hx .                  # editor
# Right: terminal              # go run, go test, git

# Or from helix: Space t       # quick terminal below
```

### Workflow 3: AI-Assisted Dev (tdl)

```bash
t                              # start tmux
tdl cx                         # dev layout with claude
# Layout:
# ┌──────────────┬─────────┐
# │              │         │
# │  helix       │ claude  │
# │              │         │
# ├──────────────┴─────────┤
# │  terminal              │
# └────────────────────────┘
```

### Workflow 4: Multi-Project

```bash
t                              # start tmux
tdlm cx                        # one tdl per subdirectory
# Switch: Alt-Left/Right       # between project windows
# Switch: Alt-Up/Down          # between sessions
```

### Workflow 5: Code Review

```bash
hx main.go handler.go model.go   # open multiple files
Space b                           # fuzzy switch buffers
Ctrl-w v                          # side-by-side view
ga                                # toggle between last 2 files
```

---

## Essential CLI Tools

### Already Installed on Your System

| Tool | What | Usage |
|------|------|-------|
| `eza` | Better `ls` | `ls` (aliased), `lt` (tree view), `lsa` (show hidden) |
| `bat` | Better `cat` | `bat file.go` — syntax highlighting + line numbers |
| `fd` | Better `find` | `fd main` — fast file search, respects .gitignore |
| `fzf` | Fuzzy finder | `ff` (aliased) — fuzzy find + preview with bat |
| `rg` | Better `grep` | `rg "func main"` — fast regex search |
| `zoxide` | Smart `cd` | `cd project` — learns your frequent dirs |
| `dust` | Better `du` | `dust` — visual disk usage |
| `tldr` | Better `man` | `tldr tar` — practical examples |
| `jq` | JSON processor | `cat data.json \| jq '.name'` |
| `lazygit` | Git TUI | `lazygit` or `Space g s` in helix |
| `lazydocker` | Docker TUI | `lazydocker` or `Super+Shift+D` |
| `btop` | System monitor | `btop` or `Super+Ctrl+T` |
| `mise` | Tool versioner | Manages Go, Node, Python versions |
| `starship` | Shell prompt | Already configured in your shell |
| `eff` | Edit fuzzy file | `eff` — find file with fzf, open in editor |

### Recommended to Install

These are the tools that will make a real difference long-term:

```bash
yay -S git-delta sd entr tokei hyperfine trash-cli ncdu atuin
```

| Tool | What | Why You Need It |
|------|------|----------------|
| **`delta`** | Git diff viewer | Syntax-highlighted, side-by-side diffs. Makes `git diff` beautiful. |
| **`sd`** | Better `sed` | `sd 'old' 'new' file.go` — intuitive find-replace, no escaping hell |
| **`entr`** | File watcher | `ls *.go \| entr go test ./...` — re-runs command when files change |
| **`tokei`** | Code counter | `tokei` — lines of code by language, fast |
| **`hyperfine`** | Benchmarker | `hyperfine 'go test' 'go test -short'` — benchmark any command |
| **`trash-cli`** | Safe delete | `trash-put file` instead of `rm` — recoverable! |
| **`ncdu`** | Disk analyzer | `ncdu /` — interactive, find what's eating your disk |
| **`atuin`** | Shell history | Synced, searchable shell history across sessions. Game changer. |

### How Each Tool Fits Into Your Workflow

#### `delta` — Beautiful Git Diffs
```bash
# Add to ~/.gitconfig (one-time setup):
[core]
    pager = delta
[interactive]
    diffFilter = delta --color-only
[delta]
    navigate = true
    side-by-side = true
    line-numbers = true
```
Now every `git diff`, `git log -p`, `git show` looks amazing.

#### `sd` — Find and Replace Without Pain
```bash
# Replace in single file
sd 'oldFunc' 'newFunc' main.go

# Replace across all Go files
fd -e go | xargs sd 'oldFunc' 'newFunc'

# No regex escaping needed (unlike sed)
sd 'http://localhost:3000' 'https://api.example.com' config.go
```

#### `entr` — Auto-Run on File Changes
```bash
# Re-run tests when any Go file changes
fd -e go | entr -c go test ./...

# Rebuild on changes
fd -e go | entr -c go build ./...

# Restart server on changes
fd -e go | entr -r go run .

# Run in a tmux pane next to your editor — instant feedback loop!
```

#### `atuin` — Never Lose a Command Again
```bash
# After install, init in your shell:
# For zsh — add to ~/.zshrc:
eval "$(atuin init zsh)"

# For bash — add to ~/.bashrc:
eval "$(atuin init bash)"

# Then Ctrl-r gives you fuzzy-searchable history
# across all terminals, all sessions, all time
# It replaces the default Ctrl-r history search
```

#### `tokei` — Know Your Codebase
```bash
tokei                    # stats for current project
tokei src/               # stats for a directory
# Shows lines of code, comments, blanks per language
```

#### Combining Tools (Power Combos)

```bash
# Find all TODO comments in Go files
rg 'TODO|FIXME|HACK' --type go

# Find large files eating disk
fd --size +10m

# Find recently modified Go files
fd -e go --changed-within 1d

# Interactive file opener
fd -e go | fzf --preview 'bat --color=always {}' | xargs hx

# Benchmark two approaches
hyperfine 'go test -run TestOld' 'go test -run TestNew'

# Find and replace across entire project safely
fd -e go | xargs sd 'OldStruct' 'NewStruct'
# Then: git diff to verify, git commit if good

# Watch tests while coding (in tmux pane)
fd -e go | entr -c go test ./... 2>&1 | head -20
```

---

## Your Config Files

| File | Purpose |
|------|---------|
| `~/.config/helix/config.toml` | Helix editor settings + keybindings |
| `~/.config/helix/languages.toml` | LSP and language config |
| `~/.config/tmux/tmux.conf` | Tmux settings |
| `~/.config/hypr/bindings.conf` | Your Hyprland keybindings |
| `~/.config/starship.toml` | Shell prompt |
| `~/.bashrc` | Bash config (Omarchy default, aliases, exports) |
| `~/.zshrc` | Zsh config (interactive shell — aliases, tools, prompt) |

---

## LSP Setup

Your helix has language servers configured for:

| Language | Server | Status |
|----------|--------|--------|
| Go | gopls | ✓ Installed + configured |
| Go (debug) | dlv | ✓ Installed |
| Go (format) | goimports | ✓ Installed |
| Python | pyright/pylsp | Install: `yay -S pyright` |
| Rust | rust-analyzer | Install: `rustup component add rust-analyzer` |
| TypeScript | typescript-language-server | Install: `npm i -g typescript-language-server typescript` |
| JavaScript | Same as TypeScript | Same install |
| Bash | bash-language-server | Install: `npm i -g bash-language-server` |
| TOML | taplo | Install: `yay -S taplo` |
| YAML | yaml-language-server | Install: `npm i -g yaml-language-server` |
| Markdown | marksman | Install: `yay -S marksman` |

Install what you need:
```bash
# Python
yay -S pyright

# Rust
rustup component add rust-analyzer

# Web (TypeScript/JavaScript)
npm i -g typescript-language-server typescript

# Config files
npm i -g bash-language-server yaml-language-server
yay -S taplo marksman
```

Verify with:
```bash
helix --health go
helix --health python
helix --health rust
helix --health typescript
```

---

## Zsh on Omarchy

Omarchy ships with bash, but zsh works great as your interactive shell.

### How It Works

- **Login shell**: Your login shell is set to zsh (`chsh -s /usr/bin/zsh`)
- **Terminal**: Ghostty is configured to launch zsh (`command = /usr/bin/zsh` in `~/.config/ghostty/config`)
- **Boot chain**: Hyprland launches via `uwsm` (a compiled binary) — it doesn't depend on bash as your login shell
- **Omarchy internals**: The `~/.config/uwsm/env` file handles session environment (PATH, mise) independently of your shell

### What NOT to Do

| Don't | Why |
|-------|-----|
| `source ~/.local/share/omarchy/default/bash/rc` in `.zshrc` | Runs `starship init bash` which outputs `\[\]` escape garbage in zsh |
| `source ~/.local/share/omarchy/default/bash/shell` in `.zshrc` | Uses `shopt` (bash-only builtin), will error |
| `source ~/.local/share/omarchy/default/bash/init` in `.zshrc` | Runs `bind -f` (bash-only), inits tools in bash mode |

### What TO Do

Cherry-pick the compatible parts and init tools in zsh mode:

```zsh
# Omarchy environment (same vars as bash)
export OMARCHY_PATH=$HOME/.local/share/omarchy
export PATH=$OMARCHY_PATH/bin:$PATH:$HOME/.local/bin
export BAT_THEME=ansi

# Omarchy aliases & functions (zsh-compatible, safe to source)
source ~/.local/share/omarchy/default/bash/aliases 2>/dev/null
for f in $OMARCHY_PATH/default/bash/fns/*; do source "$f" 2>/dev/null; done

# Init tools in ZSH mode (not bash mode!)
eval "$(starship init zsh)"
eval "$(zoxide init zsh)"
eval "$(mise activate zsh)"
source <(fzf --zsh)
```

### Zsh-Specific Features You Get

| Feature | What It Does |
|---------|--------------|
| `setopt AUTO_CD` | Type a directory name to `cd` into it (no `cd` needed) |
| `setopt CORRECT` | Suggests spelling fixes for mistyped commands |
| `setopt SHARE_HISTORY` | History shared across all terminal sessions in real-time |
| Tab completion menus | `autoload -Uz compinit && compinit` — interactive menu completion |
| `bindkey '^[[A'` | Up arrow searches history matching what you've typed so far |
| Glob qualifiers | `ls *(.)` (files only), `ls *(/)` (dirs only), `ls *(.om[1,5])` (5 newest files) |
| `**/*.go` | Recursive globbing built-in (no `find` needed) |

### Useful Zsh Tricks

```zsh
# Glob qualifiers — filter files without find
ls **/*.go              # all Go files recursively
ls *(.)                 # only files (not dirs) in current dir
ls *(/)                 # only directories
ls *(.om[1,5])          # 5 most recently modified files
ls *(.Lm+10)            # files larger than 10MB

# Quick substitution in last command
# Ran: go test -run TestOld
^Old^New                # reruns as: go test -run TestNew

# Expand aliases before running (to see what they do)
# Type the alias, then press Ctrl-x a
ls<Ctrl-x a>            # expands to: eza -lh --group-directories-first --icons=auto

# Suffix aliases — open files by extension
alias -s go=hx          # typing "main.go" opens it in helix
alias -s md=bat          # typing "README.md" runs bat on it

# Global aliases — expand anywhere in a command
alias -g G='| rg'        # ls G pattern  →  ls | rg pattern
alias -g L='| less'      # cat file G err  →  cat file | rg err
alias -g C='| wl-copy'   # echo foo C  →  echo foo | wl-copy
```

---

## Omarchy-Specific Gotchas

| Problem | Cause | Fix |
|---------|-------|-----|
| `Super+W` closes window | Hyprland keybinding | Use `:w` / `Ctrl-s` / `Space w` to save |
| `Alt+1-9` doesn't work in helix | Tmux intercepts for windows | Use `Space b` or `gn/gp` for buffers |
| Clipboard doesn't paste in helix | Different clipboard path | Use `Space y` / `Space p` (custom bindings) |
| Theme looks wrong | Terminal issue | `true-color = true` already set |
| `n` alias conflicts | Omarchy aliases `n` to nvim | Override in .bashrc/.zshrc if needed |
| `\[\]` garbage in zsh prompt | Sourcing Omarchy's `bash/rc` in zsh | Don't source it — cherry-pick aliases + init tools in zsh mode |

---

## Learning Path

| Week | Focus | Practice |
|------|-------|----------|
| 1 | Movement + insert | `hjkl`, `wb`, `iao`, `:w`, `jk` to escape |
| 2 | Selection + actions | `v`, `x`, `d/c/y/p`, `/` search |
| 3 | Text objects + space | `mi/ma`, Space menu, `gd/gr` |
| 4 | Multi-cursor + surround | `C`, `s`, `ms/md/mr`, macros |
| 5 | LSP + tmux workflows | `tdl`, splits, `entr`, full dev cycle |

Start practicing:
```bash
hx ~/helix-tutor/tutor.go
```
