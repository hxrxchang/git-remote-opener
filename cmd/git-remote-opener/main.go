package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hxrxchang/git-remote-opener/gitrepo"
	"github.com/skratchdot/open-golang/open"
)

type ICommander interface {
	GetGitRemoteInfo() ([]byte, error)
	Printf(msg string)
	PrintErr(msg error)
}

type Commander struct{}

func (c *Commander) GetGitRemoteInfo() ([]byte, error) {
	out, err := exec.Command("git", "remote", "-v").CombinedOutput()
	return out, err
}

func (c *Commander) Printf(msg string) {
	fmt.Println(msg)
}

func (c *Commander) PrintErr(msg error) {
	fmt.Println(msg)
}

func execute() int {
	var commander ICommander = &Commander{}
	out, err := commander.GetGitRemoteInfo()
	if err != nil {
		msg := "fatal: not a git repository (or any of the parent directories): .git"
		commander.Printf(msg)
		return 1
	}

	stringified := string(out)
	if stringified == "" {
		msg := "fatal: 'origin' does not appear to be a git repository\nfatal: Could not read from remote repository.\n\nPlease make sure you have the correct access rights\nand the repository exists."
		commander.Printf(msg)
		return 1
	}

	replaced := strings.Replace(stringified, `\n`, "\n", -1)
	splited := strings.Split(replaced, "\n")
	origin := splited[0]
	originURL, err := gitrepo.GetRepoURL(origin)
	if err != nil {
		commander.PrintErr((err))
		return 1
	}

	error := open.Run(originURL)
	if err != nil {
		commander.PrintErr(error)
		return 1
	}

	return 0
}

func main() {
	os.Exit(execute())
}
