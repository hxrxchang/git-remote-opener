package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Open the repository
	r, err := git.PlainOpen(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open repository: %v\n", err)
		os.Exit(1)
	}

	// Get the remote information
	remotes, err := r.Remotes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get remote information: %v\n", err)
		os.Exit(1)
	}

	// Print the remote information
	for _, remote := range remotes {
		fmt.Printf("Name: %s\n", remote.Config().Name)
		fmt.Printf("URL: %s\n", remote.Config().URLs[0])
	}
}
