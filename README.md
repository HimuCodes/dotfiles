# Dotfiles

Personal dotfiles for Omarchy (Arch Linux + Hyprland), managed with [GNU Stow](https://www.gnu.org/software/stow/).

## One-Command Setup

```bash
curl -fsSL https://raw.githubusercontent.com/HimuCodes/samsara/main/bootstrap.sh | bash
```

Or see [BOOTSTRAP.md](BOOTSTRAP.md) for the full walkthrough.

## Packages

| Package | What it manages |
|---------|----------------|
| `bash` | `~/.bashrc` — shell config, aliases, exports |
| `zsh` | `~/.zshrc` — zsh config, prompt, history, completions |
| `helix` | `~/.config/helix/` — editor config + language servers |
| `tmux` | `~/.config/tmux/tmux.conf` — terminal multiplexer |
| `starship` | `~/.config/starship.toml` — cross-shell prompt |
| `hypr` | `~/.config/hypr/` — Hyprland WM, bindings, shaders |
| `ssh` | `~/.ssh/config` — SSH client config, 1Password agent |
| `helix-tutor` | `~/helix-tutor/` — Helix editor tutorial + GUIDE.md |

## Quick Start

### Fresh Machine

```bash
# Clone samsara and run bootstrap (handles everything)
git clone git@github.com:HimuCodes/samsara.git ~/samsara
~/samsara/bootstrap.sh
```

### Already Have Dotfiles Cloned

```bash
cd ~/dotfiles
./bootstrap.sh
```

### Manual Stow (No Bootstrap)

```bash
yay -S stow
git clone git@github.com:HimuCodes/dotfiles.git ~/dotfiles
cd ~/dotfiles
stow bash zsh helix tmux starship hypr ssh helix-tutor
```

## How Stow Works

Stow creates **symlinks** from your home directory into the dotfiles repo. Each "package" is a directory inside `~/dotfiles/` that mirrors the structure of your home directory.

```
~/dotfiles/helix/.config/helix/config.toml
    ↓ stow creates symlink ↓
~/.config/helix/config.toml → ~/dotfiles/helix/.config/helix/config.toml
```

You edit the files in `~/dotfiles/` (or follow the symlinks — same file either way), commit, and push.

## Day-to-Day

```bash
# Edit any config (symlink means you edit the dotfiles repo directly)
hx ~/.config/helix/config.toml

# Commit changes
cd ~/dotfiles
git add -A
git commit -m "helix: add markdown format-on-save"
git push
```

## Adding a New Config

```bash
mkdir -p ~/dotfiles/alacritty/.config/alacritty
mv ~/.config/alacritty/alacritty.toml ~/dotfiles/alacritty/.config/alacritty/
cd ~/dotfiles
stow alacritty
git add -A && git commit -m "add alacritty config"
```

## Removing a Package

```bash
cd ~/dotfiles
stow -D <package>   # removes symlinks, doesn't delete files
```

## Resolving Conflicts

If stow says "existing target is not owned by stow":

```bash
# Option A: adopt then restore
cd ~/dotfiles && stow --adopt helix && git checkout -- .

# Option B: back up and retry
mv ~/.config/helix/config.toml ~/.config/helix/config.toml.bak
cd ~/dotfiles && stow helix
```

## System Info

- **OS**: Omarchy (Arch Linux)
- **WM**: Hyprland
- **Terminal**: Ghostty
- **Shell**: zsh (with bash fallback for Omarchy internals)
- **Editor**: Helix
- **Multiplexer**: tmux (prefix: Ctrl-Space)
- **Prompt**: Starship
- **Runtimes**: mise (Go, Node, Python) + rustup (Rust)
- **SSH**: 1Password agent

## Related

- [HimuCodes/samsara](https://github.com/HimuCodes/samsara) — standalone bootstrap entry point
- [BOOTSTRAP.md](BOOTSTRAP.md) — detailed bootstrap documentation
