package git_repo

import "regexp"

type RepoInfo struct {
	Hostname, User, Repo string
}

func GetRepoUrl(remoteUrlString string) string {
	whenSshRegexp := regexp.MustCompile(`^origin\s+git@`)
	whenHttpsRegexp := regexp.MustCompile(`^origin\s+https:`)
	var repoInfo RepoInfo
	if whenSshRegexp.MatchString(remoteUrlString) {
		repoInfo = buildRepoInfo(remoteUrlString, `^origin\s+git@(.*):(.*)\/(.*).git`)
	} else if whenHttpsRegexp.MatchString(remoteUrlString) {
		repoInfo = buildRepoInfo(remoteUrlString, `^origin\s+https:\/\/(.*)\/(.*)\/(.*).git`)
	}
	return "https://" + repoInfo.Hostname + "/" + repoInfo.User + "/" + repoInfo.Repo
}

func buildRepoInfo(remoteUrlString string, regexpString string) RepoInfo {
	regexp := regexp.MustCompile(regexpString)
	matchingResult := regexp.FindStringSubmatch(remoteUrlString)
	return RepoInfo{matchingResult[1], matchingResult[2], matchingResult[3]}
}
