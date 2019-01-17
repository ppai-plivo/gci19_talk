package main

import (
	"expvar"
	"fmt"
)

// START OMIT

// go build -ldflags "-X main.built=$(date +%FT%T%z) -X main.gitSHA=`git rev-parse --short HEAD`"

var (
	gitSHA = "undefined"
	built  = "undefined"
)

var (
	expGitSHA = expvar.NewString("gitSHA")
	expBuilt  = expvar.NewString("built")
)

func init() {
	expGitSHA.Set(gitSHA)
	expBuilt.Set(built)
}

// END OMIT
