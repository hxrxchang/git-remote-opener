package git_repo

import "regexp"

func GetRepoUrl(str string) string {
	// put "?:" not to capture
	reg := regexp.MustCompile(`(?:git@github.com:|https:\/\/github.com\/)(.*).git`)
	s := reg.FindStringSubmatch(str)
	// get value of s[1] to get captured value by "(.*)"
	return "https://github.com/" + s[1]
}
