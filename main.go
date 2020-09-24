package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hxrxchang/git-remote-opener/git_repo"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
		os.Exit(1)
		return
	}
	stringified := string(out)
	if stringified == "" {
		fmt.Println("fatal: 'origin' does not appear to be a git repository\nfatal: Could not read from remote repository.\n\nPlease make sure you have the correct access rights\nand the repository exists.")
		os.Exit(1)
		return
	}
	replaced := strings.Replace(stringified, `\n`, "\n", -1)
	splited := strings.Split(replaced, "\n")
	origin := splited[0]
	originURL, err := git_repo.GetRepoUrl(origin)
	if err != nil {
		fmt.Println(err)
		return
	}
	open.Run(originURL)
}
