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

source ~/.local/share/omarchy/default/bash/rc 2>/dev/null
