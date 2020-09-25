package gitrepo

import (
	"errors"
	"regexp"
)

type repoInfo struct {
	Hostname, User, Repo string
}

// GetRepoURL is a function to get string which means git remote repository's url and is started with https://.+
func GetRepoURL(remoteURLString string) (string, error) {
	whenSSHRegexp := regexp.MustCompile(`^origin\s+git@`)
	whenHTTPSRegexp := regexp.MustCompile(`^origin\s+https:`)
	var repoInfoVal repoInfo
	var repoInfoErr error
	if whenSSHRegexp.MatchString(remoteURLString) {
		repoInfoVal, repoInfoErr = buildRepoInfo(remoteURLString, `^origin\s+git@(.*):(.*)\/(.*).git`)
	} else if whenHTTPSRegexp.MatchString(remoteURLString) {
		repoInfoVal, repoInfoErr = buildRepoInfo(remoteURLString, `^origin\s+https:\/\/(.*)\/(.*)\/(.*).git`)
	} else {
		return "", errors.New("something went wrong")
	}
	if repoInfoErr != nil {
		return "", errors.New("something went wrong")
	}
	return "https://" + repoInfoVal.Hostname + "/" + repoInfoVal.User + "/" + repoInfoVal.Repo, nil
}

func buildRepoInfo(remoteURLString string, regexpString string) (repoInfo, error) {
	repoInfoRegexp := regexp.MustCompile(regexpString)
	matchingResult := repoInfoRegexp.FindStringSubmatch(remoteURLString)
	if len(matchingResult) == 0 {
		return repoInfo{"", "", ""}, errors.New("something went wrong")
	}
	return repoInfo{matchingResult[1], matchingResult[2], matchingResult[3]}, nil
}
