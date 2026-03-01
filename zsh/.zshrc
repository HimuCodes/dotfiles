export EDITOR=helix
export VISUAL=helix
export SUDO_EDITOR=helix
export SSH_AUTH_SOCK=~/.1password/agent.sock
alias hx='helix'

. "$HOME/.local/share/../bin/env"

autoload -Uz compinit && compinit
autoload -Uz vcs_info
precmd() { vcs_info }
setopt PROMPT_SUBST

zstyle ':vcs_info:git:*' formats ' %F{cyan}(%b)%f'

PROMPT='%F{blue}%~%f${vcs_info_msg_0_} %F{white}❯%f '

setopt AUTO_CD
setopt CORRECT
setopt HIST_IGNORE_DUPS
setopt HIST_IGNORE_ALL_DUPS
setopt HIST_FIND_NO_DUPS
setopt SHARE_HISTORY
setopt INC_APPEND_HISTORY

HISTFILE=~/.zsh_history
HISTSIZE=50000
SAVEHIST=50000

bindkey -e
bindkey '^[[A' history-search-backward
bindkey '^[[B' history-search-forward

# Omarchy environment
export OMARCHY_PATH=$HOME/.local/share/omarchy
export PATH=$OMARCHY_PATH/bin:$PATH:$HOME/.local/bin
export BAT_THEME=ansi

# Omarchy aliases & functions (zsh-compatible)
source ~/.local/share/omarchy/default/bash/aliases 2>/dev/null
for f in $OMARCHY_PATH/default/bash/fns/*; do source "$f" 2>/dev/null; done

# Tool initialization (zsh mode)
if command -v starship &> /dev/null; then
  eval "$(starship init zsh)"
fi

if command -v zoxide &> /dev/null; then
  eval "$(zoxide init zsh)"
fi

if command -v mise &> /dev/null; then
  eval "$(mise activate zsh)"
fi

if command -v fzf &> /dev/null; then
  source <(fzf --zsh) 2>/dev/null
fi
