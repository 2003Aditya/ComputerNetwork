!/bin/bash

tmux new-session -d -s comp

#first window
tmux send-keys "nvim ~/Code/ComputerNetwork" C-m
tmux rename-window "Code"

#second window
tmux new-window -t comp:2 -n "sender"
tmux send-keys "nvim ~/Code/ComputerNetwork -c 'terminal'" C-m

#third window
tmux new-window -t comp:3 -n "receiver"
tmux send-keys "nvim ~/Code/ComputerNetwork -c 'terminal'" C-m



tmux attach -t comp

