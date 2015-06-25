#!/bin/bash
echo "Starting new tmux session: main"
PI_HOME=/home/pi
APNAME=`cat $PI_HOME/AIRPLAY_NAME`
HWADDR=`python $PI_HOME/home-audio/gen_hwaddr.py`
echo "apname: $APNAME"
echo "hwaddr: $HWADDR"
tmux new -d -s main
tmux send-keys -t main "cd /home/pi/shairplay" ENTER
tmux send-keys -t main "/usr/local/bin/shairplay --apname=\"$APNAME\" --password=\"\" --hwaddr=\"$HWADDR\"" ENTER
