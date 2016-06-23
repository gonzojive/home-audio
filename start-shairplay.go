// Program start-shairplay starts a Shairplay server and restarts it if it
// crashes.
package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"time"
)

const (
	airplayNameFile    = "/home/pi/AIRPLAY_NAME"
	sleepAfterFailTime = time.Second * 5
)

func genHardwareAddr(r *rand.Rand) net.HardwareAddr {
	b := make([]byte, 6)
	for i := 0; i < len(b); i++ {
		b[i] = byte(r.Int31n(256))
	}
	return net.HardwareAddr(b)
}

func randFromName(s string) *rand.Rand {
	return rand.New(rand.NewSource(int64(crc32.ChecksumIEEE([]byte(s)))))
}

func runCmdOnce() error {
	nameBytes, err := ioutil.ReadFile(airplayNameFile)
	if err != nil {
		log.Fatalf("Error reading Airplay device name: %v", err)
	}
	name := string(nameBytes)

	// Note: Consider setting --ao_devicename="iec958:DAC", where
	// "iec958:DAC" is obtained from `aplay -L`
	args := []string{
		fmt.Sprintf("--apname=%s", name),
		"--password=",
		fmt.Sprintf("--hwaddr=%s", genHardwareAddr(randFromName(name))),
		"--ao_devicename=iec958:DAC",
	}
	cmd := exec.Command("shairplay", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Printf("Starting shairplay with args: %q", args)
	return cmd.Run()
}

func main() {
	for {
		if err := runCmdOnce(); err != nil {
			log.Printf("Error during run: %v", err)
			log.Printf("Sleeping for %v before restarting...", sleepAfterFailTime)
			time.Sleep(sleepAfterFailTime)
		} else {
			break
		}
	}
}
