package gitrepo

import (
	"testing"
)

func TestGetRepoURLWithSshString(t *testing.T) {
	result, _ := GetRepoURL("origin  git@github.com:hxrxchang/git-remote-opener.git (fetch)")
	if result != "https://github.com/hxrxchang/git-remote-opener" {
		t.Fatal("test failed")
	}
}

func TestGetRepoURLWithHttpsString(t *testing.T) {
	result, _ := GetRepoURL("origin	https://github.com/hxrxchang/git-remote-opener.git (fetch)")
	if result != "https://github.com/hxrxchang/git-remote-opener" {
		t.Fatal("test failed")
	}
}

func TestGetRepoURLWithInvalidString(t *testing.T) {
	result, _ := GetRepoURL("invalid string")
	if result != "" {
		t.Fatal("test failed")
	}
}
