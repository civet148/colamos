package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"colamos/app"
	"colamos/cmd/colamosd/cmd"
)

const (
	Version = "v0.1.0"
)

var (
	BuildTime = ""
	GitCommit = ""
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	version.Version = fmt.Sprintf("%s %s %s", Version, BuildTime, GitCommit)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
