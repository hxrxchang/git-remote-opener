package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/hxrxchang/git-remote-opener/v3/cmd/git-remote-open/mock"
)

func Test_Main(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommander := mock_main.NewMockICommander(ctrl)
	out := []byte("origin  git@github.com:hxrxchang/git-remote-opener.git (fetch)\norigin  git@github.com:hxrxchang/git-remote-opener.git (push)")
	mockCommander.EXPECT().GetGitRemoteInfo().Return(out, nil)
	mockCommander.EXPECT().Open("https://github.com/hxrxchang/git-remote-opener").Return(nil)

	result := _main(mockCommander)

	if result != 0 {
		t.Fatal("result must be 0")
	}
}

func Test_MainWhenNotGitRepo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommander := mock_main.NewMockICommander(ctrl)
	msg := "fatal: not a git repository (or any of the parent directories): .git"
	out := []byte(msg)
	err := errors.New("exit status 128")
	mockCommander.EXPECT().GetGitRemoteInfo().Return(out, err)
	mockCommander.EXPECT().Println(msg)

	result := _main(mockCommander)

	if result != 1 {
		t.Fatal("result must be 1")
	}
}

func Test_MainWhenWithoutGitRemote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommander := mock_main.NewMockICommander(ctrl)
	out := []byte("")
	mockCommander.EXPECT().GetGitRemoteInfo().Return(out, nil)
	mockCommander.EXPECT().Println("The remote repository is not configured.")

	result := _main(mockCommander)

	if result != 1 {
		t.Fatal("result must be 1")
	}
}
