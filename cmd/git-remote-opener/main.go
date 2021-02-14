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
}

type Commander struct{}

func (c *Commander) GetGitRemoteInfo() ([]byte, error) {
	out, err := exec.Command("git", "remote", "-v").CombinedOutput()
	return out, err
}

func execute() int {
	var commander ICommander = &Commander{}
	out, err := commander.GetGitRemoteInfo()
	if err != nil {
		fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
		return 1
	}

	stringified := string(out)
	if stringified == "" {
		fmt.Println("fatal: 'origin' does not appear to be a git repository\nfatal: Could not read from remote repository.\n\nPlease make sure you have the correct access rights\nand the repository exists.")
		return 1
	}

	replaced := strings.Replace(stringified, `\n`, "\n", -1)
	splited := strings.Split(replaced, "\n")
	origin := splited[0]
	originURL, err := gitrepo.GetRepoURL(origin)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	error := open.Run(originURL)
	if err != nil {
		fmt.Println(error)
		return 1
	}

	return 0
}

func main() {
	os.Exit(execute())
}
