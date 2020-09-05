package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	out, _ := exec.Command("git", "remote", "-v").Output()
	stringified := string(out)
	replaced := strings.Replace(stringified, `\n`, "\n", -1)
	splited := strings.Split(replaced, "\n")
	origin := splited[0]
	origin = "https://github.com/" + getOrigin(origin)
	fmt.Println(origin)
}

func getOrigin(str string) string {
	reg := regexp.MustCompile(`:(.*).git`)
	s := reg.FindStringSubmatch(str)
	return s[1]
}
