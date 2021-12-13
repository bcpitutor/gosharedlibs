package gosharedlibs

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"runtime"
)

func InitalizeTikiFolder() {

	home, err := HomeFolder()
	if err != nil {
		fmt.Println(err)
	}
	outdirpath := fmt.Sprintf("%s%s", home, "/.tikitool")

	fMode := fs.FileMode(uint32(0700))
	err = os.MkdirAll(outdirpath, fs.FileMode(fMode))
	if err != nil {
		fmt.Println(err)
	}

}

func HomeFolder() (string, error) {

	switch runtime.GOOS {

	case "darwin":
		home := os.Getenv("HOME")
		if home == "" {
			fmt.Println("Home folder couldn't found. Please set your home folder in your shell\n$ export HOME=/Users/$(whoami)")
			return "", errors.New("home couldn't found")
		}
		return home, nil
	case "linux":
		home := os.Getenv("HOME")
		if home == "" {
			fmt.Println("Home folder couldn't found. Please set your home folder in your shell\n$ export HOME=/home/$(whoami) ")
			return "", errors.New("home couldn't found")
		}
		return home, nil

	case "windows":
		home := os.Getenv("HOME")
		if home == "" {
			fmt.Println("Home folder couldn't found. Please set your home folder in your shell\n$ export HOME=/Users/$(whoami)")
			return "", errors.New("home couldn't found")
		}

		drive := os.Getenv("HOMEDRIVE")
		path := os.Getenv("HOMEPATH")
		home = drive + path
		if drive == "" || path == "" {
			return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")

		}

		return home, nil

	default:
		return "", errors.New("home couldn't found")
	}
}
