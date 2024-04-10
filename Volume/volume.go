package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rander *rand.Rand

var maxVolume int
var minVolume int

func init() {
	rander = rand.New(rand.NewSource(time.Now().Unix()))
	flag.IntVar(&maxVolume, "max", 40, "max volume")
	flag.IntVar(&minVolume, "min", 30, "min volume")
}
func main() {
	flag.Parse()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("MAX:", maxVolume, " min:", minVolume)
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			VolumeSetter(rander.Intn(maxVolume-minVolume) + minVolume)
		case <-sigs:
			ticker.Stop()
			fmt.Println("programe exit")
			return
		}
	}
}
