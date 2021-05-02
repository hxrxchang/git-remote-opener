![CI](https://github.com/hxrxchang/git-remote-opener/workflows/CI/badge.svg)

# git-remote-opener

git-remote-opener is a command line tool to open a web page of a remote repository.

The following git hosting services are operation checked.

- GitHub
- GitLab
- Bitbucket

## Installation

```
$ go get github.com/hxrxchang/git-remote-opener/v2/cmd/gro
```

If you don't use Go, download the binary from [GitHub Releases](https://github.com/hxrxchang/git-remote-opener/releases) and drop it in your \$PATH.  
(For example, `/usr/local/bin` or `$HOME/bin`)

## Usage

```
$ cd <path-to-your-project>
$ gro
```

Then, browser will open a remote repository's page of your project.
