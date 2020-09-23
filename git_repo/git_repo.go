package git_repo

import (
	"errors"
	"regexp"
)

type RepoInfo struct {
	Hostname, User, Repo string
}

func GetRepoUrl(remoteUrlString string) (string, error) {
	whenSshRegexp := regexp.MustCompile(`^origin\s+git@`)
	whenHttpsRegexp := regexp.MustCompile(`^origin\s+https:`)
	var repoInfo RepoInfo
	var repoInfoErr error
	if whenSshRegexp.MatchString(remoteUrlString) {
		repoInfo, repoInfoErr = buildRepoInfo(remoteUrlString, `^origin\s+git@(.*):(.*)\/(.*).git`)
	} else if whenHttpsRegexp.MatchString(remoteUrlString) {
		repoInfo, repoInfoErr = buildRepoInfo(remoteUrlString, `^origin\s+https:\/\/(.*)\/(.*)\/(.*).git`)
	} else {
		return "", errors.New("something went wrong")
	}
	if repoInfoErr != nil {
		return "", errors.New("something went wrong")
	}
	return "https://" + repoInfo.Hostname + "/" + repoInfo.User + "/" + repoInfo.Repo, nil
}

func buildRepoInfo(remoteUrlString string, regexpString string) (RepoInfo, error) {
	repoInfoRegexp := regexp.MustCompile(regexpString)
	matchingResult := repoInfoRegexp.FindStringSubmatch(remoteUrlString)
	if len(matchingResult) == 0 {
		return RepoInfo{"", "", ""}, errors.New("something went wrong")
	}
	return RepoInfo{matchingResult[1], matchingResult[2], matchingResult[3]}, nil
}
