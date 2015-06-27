#!/bin/bash
echo "Starting new tmux session: main"
tmux new -d -s main

tmux send-keys -t main "cd /home/pi/shairplay" ENTER
tmux send-keys "go run ../home-audio/start-shairplay.go" ENTER
#tmux send-keys -t main "/usr/local/bin/shairplay --apname=\"$APNAME\" --password=\"\" --hwaddr=\"$HWADDR\"" ENTER
