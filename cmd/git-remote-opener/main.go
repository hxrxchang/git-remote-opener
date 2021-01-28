package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hxrxchang/git-remote-opener/gitrepo"
	"github.com/skratchdot/open-golang/open"
)

func runCommand() int {
	out, err := exec.Command("git", "remote", "-v").Output()
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
	open.Run(originURL)
	return 0
}

func main() {
	os.Exit(runCommand())
}
