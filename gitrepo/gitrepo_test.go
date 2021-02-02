package gitrepo

import (
	"testing"
)

func TestGetRepoURLWithSshString(t *testing.T) {
	result, err := GetRepoURL("origin  git@github.com:hxrxchang/git-remote-opener.git (fetch)")
	if err != nil {
		t.Error(err)
	}
	if result != "https://github.com/hxrxchang/git-remote-opener" {
		t.Fatal("must return valid url")
	}
}

func TestGetRepoURLWithHttpsString(t *testing.T) {
	result, err := GetRepoURL("origin	https://github.com/hxrxchang/git-remote-opener.git (fetch)")
	if err != nil {
		t.Error(err)
	}
	if result != "https://github.com/hxrxchang/git-remote-opener" {
		t.Fatal("must return valid url")
	}
}

func TestGetRepoURLWithInvalidString(t *testing.T) {
	result, err := GetRepoURL("invalid string")
	if err == nil {
		t.Error("must return error")
	}
	if result != "" {
		t.Fatal("must return empty string")
	}
}
