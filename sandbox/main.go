package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("git", "branch").CombinedOutput()
	stringified := string(out)
	fmt.Println(stringified)
}
