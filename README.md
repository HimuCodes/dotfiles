# Dotfiles

Personal dotfiles for Omarchy (Arch Linux + Hyprland), managed with [GNU Stow](https://www.gnu.org/software/stow/).

## Packages

| Package | What it manages |
|---------|----------------|
| `bash` | `~/.bashrc` — shell config, aliases, exports |
| `zsh` | `~/.zshrc` — zsh config, prompt, history, completions |
| `helix` | `~/.config/helix/` — editor config + language servers |
| `tmux` | `~/.config/tmux/tmux.conf` — terminal multiplexer |
| `starship` | `~/.config/starship.toml` — cross-shell prompt |
| `hypr` | `~/.config/hypr/` — Hyprland WM, bindings, shaders |
| `helix-tutor` | `~/helix-tutor/` — interactive helix learning files |

## Quick Start

### Fresh Machine Setup

```bash
# 1. Install stow
yay -S stow

# 2. Clone this repo
git clone git@github.com:YOUR_USERNAME/dotfiles.git ~/dotfiles

# 3. Stow everything
cd ~/dotfiles
stow bash zsh helix tmux starship hypr helix-tutor

# Or stow individual packages
stow helix
```

### How Stow Works

Stow creates **symlinks** from your home directory into the dotfiles repo. Each "package" is a directory inside `~/dotfiles/` that mirrors the structure of your home directory.

```
~/dotfiles/helix/.config/helix/config.toml
    ↓ stow creates symlink ↓
~/.config/helix/config.toml → ~/dotfiles/helix/.config/helix/config.toml
```

You edit the files in `~/dotfiles/` (or follow the symlinks — same file either way), commit, and push. On a new machine, clone + stow and you're done.

### Day-to-Day Workflow

```bash
# Edit any config (symlink means you edit the dotfiles repo directly)
helix ~/.config/helix/config.toml

# Commit changes
cd ~/dotfiles
git add -A
git commit -m "helix: add markdown format-on-save"
git push
```

### Adding a New Config

```bash
# Example: adding alacritty config
mkdir -p ~/dotfiles/alacritty/.config/alacritty
mv ~/.config/alacritty/alacritty.toml ~/dotfiles/alacritty/.config/alacritty/
cd ~/dotfiles
stow alacritty
git add -A && git commit -m "add alacritty config"
```

### Removing a Package

```bash
cd ~/dotfiles
stow -D helix-tutor   # removes symlinks, doesn't delete files
```

### Resolving Conflicts

If stow says "existing target is not owned by stow":

```bash
# Back up the existing file, then delete it
mv ~/.config/helix/config.toml ~/.config/helix/config.toml.bak
cd ~/dotfiles && stow helix
```

## Pushing to GitHub

```bash
# First time setup
cd ~/dotfiles
git remote add origin git@github.com:YOUR_USERNAME/dotfiles.git
git branch -M main
git add -A
git commit -m "initial dotfiles"
git push -u origin main

# If using HTTPS instead of SSH
git remote add origin https://github.com/YOUR_USERNAME/dotfiles.git
```

### GitHub SSH Key Setup (if needed)

```bash
# Generate key
ssh-keygen -t ed25519 -C "your@email.com"

# Copy public key
cat ~/.ssh/id_ed25519.pub
# → Paste at github.com/settings/ssh/new

# Test connection
ssh -T git@github.com
```

## System Info

- **OS**: Omarchy 3.4 (Arch Linux)
- **WM**: Hyprland
- **Shell**: zsh (with bash fallback)
- **Editor**: Helix
- **Terminal multiplexer**: Tmux (prefix: Ctrl-Space)
- **Prompt**: Starship
- **Tool manager**: mise (Go, Python, Node)
