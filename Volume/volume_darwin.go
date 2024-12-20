package main

import (
	"fmt"
	"os/exec"
)

func SetSysVolume(volume int) {
	volumeCmd := fmt.Sprintf("set volume output volume %d", volume)
	cmd := exec.Command("osascript", "-e", volumeCmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
