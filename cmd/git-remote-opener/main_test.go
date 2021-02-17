package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_main "github.com/hxrxchang/git-remote-opener/cmd/git-remote-opener/mock"
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
