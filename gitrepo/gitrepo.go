package gitrepo

import (
	"errors"
	"regexp"
)

// RepoInfo is a type of expressing git repository information
type RepoInfo struct {
	Hostname, User, Repo string
}

const invalidURLMessage string = "invalid git remote url.\nPlease run 'git remote -v' to check it."

// GetRepoURL is a function to get string which means git remote repository's url and is started with https://.+
func GetRepoURL(remoteURLString string) (string, error) {
	whenSSHRegexp := regexp.MustCompile(`^origin\s+git@`)
	whenHTTPSRegexp := regexp.MustCompile(`^origin\s+https:`)
	var repoInfoVal RepoInfo
	var repoInfoErr error
	if whenSSHRegexp.MatchString(remoteURLString) {
		repoInfoVal, repoInfoErr = buildRepoInfo(remoteURLString, `^origin\s+git@(.*):(.*)\/(.*).git`)
	} else if whenHTTPSRegexp.MatchString(remoteURLString) {
		repoInfoVal, repoInfoErr = buildRepoInfo(remoteURLString, `^origin\s+https:\/\/(.*)\/(.*)\/(.*).git`)
	} else {
		return "", errors.New(invalidURLMessage)
	}
	if repoInfoErr != nil {
		return "", errors.New(invalidURLMessage)
	}
	return "https://" + repoInfoVal.Hostname + "/" + repoInfoVal.User + "/" + repoInfoVal.Repo, nil
}

func buildRepoInfo(remoteURLString string, regexpString string) (RepoInfo, error) {
	repoInfoRegexp := regexp.MustCompile(regexpString)
	matchingResult := repoInfoRegexp.FindStringSubmatch(remoteURLString)
	if len(matchingResult) == 0 {
		return RepoInfo{"", "", ""}, errors.New(invalidURLMessage)
	}
	return RepoInfo{matchingResult[1], matchingResult[2], matchingResult[3]}, nil
}
