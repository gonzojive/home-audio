#!/bin/bash
echo "Starting new tmux session: main"
tmux new -d -s main
tmux send-keys -t main "cd /home/pi" ENTER
tmux send-keys -t main "/usr/local/bin/shairplay --apname=\"Brokedown Kitchen\" --password=\"\"" ENTER
