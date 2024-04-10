package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

var VolumeSetter SysVolume

func init() {
	VolumeSetter = GetVolumeFunc()
	if VolumeSetter == nil {
		panic("Unsupport sys")
	}
}

type SysVolume func(int)

func WinSysVolume(volume int) {

}

func LinuxSysVolume(volume int) {

}

func MacSysVolume(volume int) {
	volumeCmd := fmt.Sprintf("set volume output volume %d", volume)
	cmd := exec.Command("osascript", "-e", volumeCmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GetVolumeFunc() SysVolume {
	switch runtime.GOOS {
	case "linux":
		return LinuxSysVolume
	case "windows":
		return WinSysVolume
	case "darwin":
		return MacSysVolume
	default:
		return nil
	}
}
