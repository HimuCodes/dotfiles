#!/usr/bin/env bash
set -euo pipefail

# Samsara — Omarchy Dev Environment Bootstrap
# Rebuild your entire dev environment from nothing.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/HimuCodes/samsara/main/bootstrap.sh | bash
#   — or —
#   git clone git@github.com:HimuCodes/samsara.git && ./samsara/bootstrap.sh

DOTFILES_REPO="git@github.com:HimuCodes/dotfiles.git"
DOTFILES_DIR="$HOME/dotfiles"

RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
BOLD='\033[1m'
RESET='\033[0m'

step() { echo -e "\n${CYAN}${BOLD}=> $1${RESET}"; }
ok()   { echo -e "   ${GREEN}✓ $1${RESET}"; }
err()  { echo -e "   ${RED}✗ $1${RESET}"; }

# ---------------------------------------------------------------------------
# 1. Clone or update dotfiles
# ---------------------------------------------------------------------------
step "Dotfiles"

if [ -d "$DOTFILES_DIR/.git" ]; then
    ok "dotfiles already cloned — pulling latest"
    git -C "$DOTFILES_DIR" pull --ff-only || true
else
    git clone "$DOTFILES_REPO" "$DOTFILES_DIR"
    ok "cloned dotfiles"
fi

# ---------------------------------------------------------------------------
# 2. Install system packages (yay/pacman)
# ---------------------------------------------------------------------------
step "System packages"

PACKAGES=(
    # Shell & terminal
    zsh
    stow
    starship
    zoxide
    fzf
    bat
    eza
    fd
    ripgrep
    jq
    dust
    tldr

    # Editor
    helix

    # Dev tools
    git
    lazygit
    lazydocker
    tmux
    btop

    # Languages & tooling
    mise
    rustup

    # Extras
    git-delta
    sd
    entr
    tokei
    hyperfine
    trash-cli
    ncdu
)

install_pkg() {
    if command -v yay &>/dev/null; then
        yay -S --noconfirm --needed "$@"
    else
        sudo pacman -S --noconfirm --needed "$@"
    fi
}

install_pkg "${PACKAGES[@]}"
ok "system packages installed"

# ---------------------------------------------------------------------------
# 3. Stow dotfiles
# ---------------------------------------------------------------------------
step "Stowing dotfiles"

cd "$DOTFILES_DIR"

STOW_PACKAGES=(bash zsh helix tmux starship hypr ssh helix-tutor)

for pkg in "${STOW_PACKAGES[@]}"; do
    if [ -d "$pkg" ]; then
        stow --adopt -v "$pkg" 2>/dev/null || stow -v "$pkg"
        ok "stowed $pkg"
    fi
done

# Restore any adopted files to repo versions
git checkout -- . 2>/dev/null || true

# ---------------------------------------------------------------------------
# 4. Ghostty config (set zsh as shell)
# ---------------------------------------------------------------------------
step "Ghostty terminal config"

GHOSTTY_CONFIG="$HOME/.config/ghostty/config"
if [ -f "$GHOSTTY_CONFIG" ]; then
    if ! grep -q "^command = /usr/bin/zsh" "$GHOSTTY_CONFIG"; then
        echo -e "\n# Use zsh as interactive shell\ncommand = /usr/bin/zsh" >> "$GHOSTTY_CONFIG"
        ok "added zsh to ghostty config"
    else
        ok "ghostty already configured for zsh"
    fi
else
    ok "ghostty config not found (will be created by Omarchy)"
fi

# ---------------------------------------------------------------------------
# 5. Set zsh as default shell
# ---------------------------------------------------------------------------
step "Default shell"

if [ "$SHELL" != "/usr/bin/zsh" ]; then
    chsh -s /usr/bin/zsh
    ok "default shell set to zsh"
else
    ok "already using zsh"
fi

# ---------------------------------------------------------------------------
# 6. Language runtimes via mise
# ---------------------------------------------------------------------------
step "Language runtimes (mise)"

if command -v mise &>/dev/null; then
    mkdir -p "$HOME/.config/mise"
    cat > "$HOME/.config/mise/config.toml" << 'MISE'
[tools]
go = "latest"
node = "latest"
python = "latest"
MISE

    mise install --yes
    ok "go, node, python installed via mise"
else
    err "mise not found — skipping language runtimes"
fi

# ---------------------------------------------------------------------------
# 7. Rust toolchain
# ---------------------------------------------------------------------------
step "Rust toolchain"

if command -v rustup &>/dev/null; then
    rustup default stable
    rustup component add rust-analyzer
    ok "rust stable + rust-analyzer installed"
else
    err "rustup not found — skipping rust"
fi

# ---------------------------------------------------------------------------
# 8. Go tools (LSP, formatter, debugger)
# ---------------------------------------------------------------------------
step "Go tools"

install_go_tools() {
    eval "$(mise activate bash)"

    go install golang.org/x/tools/gopls@latest
    go install golang.org/x/tools/cmd/goimports@latest
    go install github.com/go-delve/delve/cmd/dlv@latest

    ok "gopls, goimports, dlv installed"
}

if command -v mise &>/dev/null; then
    install_go_tools
else
    err "mise not available — skipping go tools"
fi

# ---------------------------------------------------------------------------
# 9. 1Password SSH agent
# ---------------------------------------------------------------------------
step "1Password SSH agent"

if [ -S "$HOME/.1password/agent.sock" ]; then
    ok "1Password SSH agent already running"
else
    ok "1Password agent not found (install 1Password to enable)"
fi

# ---------------------------------------------------------------------------
# 10. Git config (delta pager)
# ---------------------------------------------------------------------------
step "Git config"

if command -v delta &>/dev/null; then
    git config --global core.pager delta
    git config --global interactive.diffFilter "delta --color-only"
    git config --global delta.navigate true
    git config --global delta.side-by-side true
    git config --global delta.line-numbers true
    ok "git configured with delta pager"
fi

if ! git config --global user.name &>/dev/null; then
    echo -e "   ${CYAN}Set your git identity:${RESET}"
    echo "   git config --global user.name \"Your Name\""
    echo "   git config --global user.email \"your@email.com\""
fi

# ---------------------------------------------------------------------------
# 11. Verify everything
# ---------------------------------------------------------------------------
step "Verification"

verify() {
    if command -v "$1" &>/dev/null; then
        ok "$1"
    else
        err "$1 not found"
    fi
}

verify zsh
verify helix
verify starship
verify tmux
verify zoxide
verify fzf
verify bat
verify eza
verify fd
verify rg
verify lazygit
verify mise
verify go
verify node
verify python3
verify rustc
verify gopls
verify goimports
verify dlv
verify delta

echo ""
step "Helix LSP health"
helix --health go 2>&1 | head -5 || true

echo ""
echo -e "${GREEN}${BOLD}Done!${RESET} Open a new terminal to start using zsh."
echo ""
echo "Quick start:"
echo "  t                    # start tmux"
echo "  tdl cx               # dev layout with claude"
echo "  hx .                 # open helix in current dir"
echo "  hx ~/helix-tutor/tutor.go   # practice helix"
