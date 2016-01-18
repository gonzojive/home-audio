# Installing home-audio on fresh Raspberry Pi

1. Set up wifi using [this guide](https://learn.adafruit.com/adafruits-raspberry-pi-lesson-3-network-setup/setting-up-wifi-with-occidentalis).

2. `sudo apt-get update && sudo apt-get install git tmux emacs-nox autoconf automake libtool libltdl-dev libao-dev libavahi-compat-libdnssd-dev avahi-daemon`

3. `git clone https://github.com/juhovh/shairplay.git`

4. Install shairplay:
  
  a. `cd shairplay && ./autogen.sh && ./configure && make`
  
  b. `sudo make install`

5. Start shairplay by running `start-shairplay.sh`

6. Install `rc.local` file so the shairplay daemon starts upon reboot.
