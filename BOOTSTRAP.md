# Bootstrap Documentation

Detailed walkthrough of what `bootstrap.sh` does, how to customize it, and how to troubleshoot.

For the standalone entry point, see [HimuCodes/samsara](https://github.com/HimuCodes/samsara).

---

## Overview

The bootstrap script takes a fresh Omarchy install to a fully configured dev environment in one run. It is **idempotent** — safe to run repeatedly.

```bash
~/dotfiles/bootstrap.sh        # if you already have dotfiles cloned
# — or —
~/samsara/bootstrap.sh          # standalone entry point (clones dotfiles first)
```

---

## Steps Explained

### Step 1: Clone Dotfiles

```bash
git clone git@github.com:HimuCodes/dotfiles.git ~/dotfiles
```

If `~/dotfiles/.git` already exists, runs `git pull --ff-only` instead. Uses SSH by default — edit `DOTFILES_REPO` at the top of the script if you need HTTPS.

### Step 2: Install System Packages

Installs packages via `yay` (falls back to `pacman` if yay isn't available). Uses `--needed` to skip packages that are already installed.

**Current package list:**

| Category | Packages |
|----------|----------|
| Shell & terminal | zsh, stow, starship, zoxide, fzf, bat, eza, fd, ripgrep, jq, dust, tldr |
| Editor | helix |
| Dev tools | git, lazygit, lazydocker, tmux, btop |
| Languages | mise, rustup |
| Extras | git-delta, sd, entr, tokei, hyperfine, trash-cli, ncdu |

**To add a package**: append it to the `PACKAGES` array in the script and re-run.

### Step 3: Stow Dotfiles

Symlinks config files from the dotfiles repo into your home directory using [GNU Stow](https://www.gnu.org/software/stow/).

```bash
stow --adopt -v <package>
git checkout -- .
```

The `--adopt` flag handles the case where Omarchy has already created default configs (e.g., `~/.bashrc`). It "adopts" the existing file into the repo, then `git checkout` restores the repo version — effectively replacing the default with yours.

**Stowed packages:** bash, zsh, helix, tmux, starship, hypr, ssh, helix-tutor

**To add a new stow package**: create a directory in `~/dotfiles/` mirroring the target path, add it to the `STOW_PACKAGES` array.

```bash
# Example: adding an alacritty config
mkdir -p ~/dotfiles/alacritty/.config/alacritty
mv ~/.config/alacritty/alacritty.toml ~/dotfiles/alacritty/.config/alacritty/
# Then add "alacritty" to STOW_PACKAGES in bootstrap.sh
```

### Step 4: Ghostty Config

Appends `command = /usr/bin/zsh` to `~/.config/ghostty/config` so Ghostty launches zsh instead of bash. Skips if already present.

Ghostty config is **not** stowed because Omarchy manages parts of it. The script only appends the zsh command line.

### Step 5: Default Shell

Runs `chsh -s /usr/bin/zsh` to set zsh as the login shell. Skips if already set.

This is safe on Omarchy — Hyprland launches via `uwsm` (a compiled binary) and does not depend on bash being the login shell.

### Step 6: Language Runtimes (mise)

Writes `~/.config/mise/config.toml` with:

```toml
[tools]
go = "latest"
node = "latest"
python = "latest"
```

Then runs `mise install --yes` to download and install all three.

**To change versions**: edit the config file and run `mise install`.

**To add a language**: add it to the `[tools]` section (e.g., `java = "21"`, `ruby = "latest"`).

### Step 7: Rust Toolchain

```bash
rustup default stable
rustup component add rust-analyzer
```

Installs the stable Rust toolchain and the rust-analyzer LSP component. Rustup itself is installed via pacman in Step 2.

### Step 8: Go Tools

Installs Go development tools into mise's Go bin directory:

| Tool | Purpose |
|------|---------|
| `gopls` | Go language server (LSP) |
| `goimports` | Auto-format + manage imports |
| `dlv` | Debugger |

These are installed via `go install`, which requires mise's Go to be active — the script handles this with `eval "$(mise activate bash)"`.

**To add more Go tools**: add `go install <package>@latest` lines in the `install_go_tools` function.

### Step 9: 1Password SSH Agent

Checks if `~/.1password/agent.sock` exists. This is informational only — install 1Password separately if you want SSH agent forwarding through it.

The SSH config in `~/dotfiles/ssh/.ssh/config` is already set up to use the 1Password agent.

### Step 10: Git Config

Configures git to use [delta](https://github.com/dandavison/delta) as the pager:

- Side-by-side diffs
- Line numbers
- Navigable output

Also prompts to set `user.name` and `user.email` if not already configured.

### Step 11: Verification

Checks that every expected tool is on the PATH:

```
✓ zsh        ✓ helix      ✓ starship   ✓ tmux
✓ zoxide     ✓ fzf        ✓ bat        ✓ eza
✓ fd         ✓ rg         ✓ lazygit    ✓ mise
✓ go         ✓ node       ✓ python3    ✓ rustc
✓ gopls      ✓ goimports  ✓ dlv        ✓ delta
```

Also runs `helix --health go` to verify LSP integration.

---

## Customization Guide

### Adding System Packages

Edit the `PACKAGES` array in `bootstrap.sh`:

```bash
PACKAGES=(
    # ... existing packages ...

    # Your additions
    neovim
    docker
    docker-compose
)
```

### Adding a New Stow Package

1. Create the directory structure:
   ```bash
   mkdir -p ~/dotfiles/myapp/.config/myapp
   cp ~/.config/myapp/config.toml ~/dotfiles/myapp/.config/myapp/
   ```

2. Add to `STOW_PACKAGES` in `bootstrap.sh`:
   ```bash
   STOW_PACKAGES=(bash zsh helix tmux starship hypr ssh helix-tutor myapp)
   ```

3. Test: `cd ~/dotfiles && stow -n -v myapp` (`-n` = dry run)

### Adding Language Runtimes

Edit `~/.config/mise/config.toml`:

```toml
[tools]
go = "latest"
node = "latest"
python = "latest"
ruby = "3.3"
java = "21"
```

Then run `mise install`.

### Changing the Dotfiles Remote

Edit the top of `bootstrap.sh` (or `samsara/bootstrap.sh`):

```bash
DOTFILES_REPO="https://github.com/YourUser/dotfiles.git"  # HTTPS
DOTFILES_REPO="git@github.com:YourUser/dotfiles.git"       # SSH
```

---

## Troubleshooting

### Stow Conflicts

**Error**: `existing target is not owned by stow`

This means a real file exists where stow wants to place a symlink.

```bash
# Option A: Let stow adopt it (then restore repo version)
cd ~/dotfiles
stow --adopt helix
git checkout -- .

# Option B: Back up and remove the conflict
mv ~/.config/helix/config.toml ~/.config/helix/config.toml.bak
cd ~/dotfiles && stow helix
```

### Package Not Found

**Error**: `target not found: some-package`

The package might be in the AUR but not in the official repos, or it may have a different name.

```bash
# Search for the correct name
yay -Ss <keyword>
```

### mise Issues

**`go: command not found` after running bootstrap**

mise needs to be activated in your shell. Check your `~/.zshrc` has:

```bash
eval "$(mise activate zsh)"
```

Open a new terminal for it to take effect.

**Wrong version installed**

```bash
mise ls           # see what's installed
mise install go@1.23  # install a specific version
mise use go@1.23      # set it as active
```

### Zsh Prompt Garbage (`\[\]`)

If you see `\[\]` in your prompt, something is running `starship init bash` in zsh. Check that `~/.zshrc` does NOT source Omarchy's `bash/rc`:

```bash
# BAD — do not put this in .zshrc
source ~/.local/share/omarchy/default/bash/rc

# GOOD — init tools in zsh mode
eval "$(starship init zsh)"
eval "$(zoxide init zsh)"
eval "$(mise activate zsh)"
```

See the "Zsh on Omarchy" section in `~/helix-tutor/GUIDE.md` for full details.

### Ghostty Not Using zsh

Check `~/.config/ghostty/config` has:

```
command = /usr/bin/zsh
```

If the line exists but is being ignored, make sure it's not overridden by an earlier `command =` line in the same file.

---

## Architecture

```
samsara/bootstrap.sh     ← standalone entry point (curl-able)
    │
    └── clones ~/dotfiles/
            │
            ├── bootstrap.sh     ← same script, also works if dotfiles already cloned
            │
            ├── bash/.bashrc
            ├── zsh/.zshrc
            ├── helix/.config/helix/
            ├── tmux/.config/tmux/
            ├── starship/.config/starship.toml
            ├── hypr/.config/hypr/
            ├── ssh/.ssh/config
            └── helix-tutor/helix-tutor/
```

The **samsara** repo exists so you can bootstrap without having dotfiles cloned yet. It's the "first thing you run" on a fresh machine. The **dotfiles** repo holds the actual configs and can also be used directly if already cloned.

Both copies of `bootstrap.sh` are identical — keep them in sync.
