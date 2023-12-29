package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	gro "github.com/hxrxchang/git-remote-opener/v3"
	"github.com/skratchdot/open-golang/open"
)

func GetGitRemoteInfo() (string, error) {
	r, err := git.PlainOpen(".")
	if err != nil {
		return "", err
	}
	remotes, err := r.Remotes()
	if err != nil {
		return "", err
	}

	if len(remotes) == 0 {
		return "", errors.New("remote repository is not configured")
	}

	return remotes[0].Config().URLs[0], err
}

func _main() int {
	remote, err := GetGitRemoteInfo()
	if err != nil {
		fmt.Printf("%s", err)
		return 1
	}

	originURL, err := gro.GetRepoURL(remote)
	if err != nil {
		fmt.Printf("%s", err)
		return 1
	}

	err = open.Run(originURL)
	if err != nil {
		fmt.Printf("%v", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(_main())
}
