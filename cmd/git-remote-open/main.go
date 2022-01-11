package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unsafe"

	gro "github.com/hxrxchang/git-remote-opener/v3"
	"github.com/skratchdot/open-golang/open"
)

type ICommander interface {
	GetGitRemoteInfo() ([]byte, error)
	Open(string) error
}

type Commander struct{}

func (c *Commander) GetGitRemoteInfo() ([]byte, error) {
	out, err := exec.Command("git", "remote", "-v").CombinedOutput()
	return out, err
}

func (c *Commander) Open(url string) error {
	err := open.Run(url)
	return err
}

func _main(commander ICommander) int {
	out, err := commander.GetGitRemoteInfo()
	if err != nil {
		fmt.Printf("%s", *(*string)(unsafe.Pointer(&out)))
		return 1
	}

	stringified := string(out)
	if stringified == "" {
		msg := "The remote repository is not configured."
		fmt.Printf("%s",msg)
		return 1
	}

	replaced := strings.Replace(stringified, `\n`, "\n", -1)
	splited := strings.Split(replaced, "\n")
	origin := splited[0]
	originURL, err := gro.GetRepoURL(origin)
	if err != nil {
		fmt.Printf("%s", err)
		return 1
	}

	error := commander.Open(originURL)
	if err != nil {
		fmt.Printf("%v", error)
		return 1
	}

	return 0
}

func main() {
	os.Exit(_main(&Commander{}))
}
